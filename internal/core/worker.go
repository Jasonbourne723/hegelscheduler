package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"hegelscheduler/internal/config"
	"hegelscheduler/internal/dto"
	"hegelscheduler/internal/queue"
	"log"
	"net/http"
	"strconv"
)

type Worker struct {
	consumer queue.Consumer
	bs       *config.BootStrap
}

func NewWorker(bs *config.BootStrap, consumer queue.Consumer) *Worker {
	return &Worker{
		consumer: consumer,
		bs:       bs,
	}
}

func (s *Worker) Start() error {
	s.consumer.Subscribe(Queue, s.Handler)
	return nil
}

func (s *Worker) Handler(data []byte) {
	var jobExecution dto.JobExectionDto
	if err := json.Unmarshal(data, &jobExecution); err != nil {
		log.Println("json unmarshall err:" + err.Error())
		return
	}
	s.SetStatus(jobExecution.JobExectionId, "Running")

	err := s.Execute(jobExecution)
	if err != nil {
		log.Println("execute err:" + err.Error())
		s.SetStatus(jobExecution.JobExectionId, "Failed")
	} else {
		s.SetStatus(jobExecution.JobExectionId, "Success")
	}

}

func (s *Worker) SetStatus(id uint64, status string) error {
	idstr := strconv.FormatUint(id, 10)

	_, err := http.Post("http://"+s.bs.Scheduler.Host+":"+s.bs.Scheduler.Port+"/JobExection/"+status+"/"+idstr, "application/json", nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Worker) Execute(jobExection dto.JobExectionDto) error {
	jsonData, err := json.Marshal(jobExection.Payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(jobExection.Method, jobExection.TargetURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	//for k, v := range jobExection.Headers {
	//	req.Header.Set(k, string(v))
	//}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New(resp.Status)
	}
}

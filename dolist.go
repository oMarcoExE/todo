package main

import "time"

type Task struct {
	Id            int
	Task_name     string
	Deadline      time.Time
	Creation_date time.Time
	Status        string
	Priority      int
}

func (t Task) addTask(name string, deadline time.Time, Priority int) {
	t.Task_name = name
	t.Deadline = deadline
	t.Priority = Priority
	t.Creation_date = time.Now()
	t.Status = "pending"
}

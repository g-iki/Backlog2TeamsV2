package backlog

import (
	model "Backlog2Teams/src/model"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/guregu/dynamo"
)

const TABLE_NAME = "BacklogTools"

var table dynamo.Table

func (m *Backlog) Create() error {
	table = model.DB.Table(TABLE_NAME)
	err := table.Put(m).If("attribute_not_exists(ID)").Run()
	return err
}

func (m *Backlog) ReadAll() []Backlog {
	list := scan()

	return list
}

func (m *Backlog) GetItem(id string) (Backlog, error) {
	b := new(Backlog)

	table = model.DB.Table(TABLE_NAME)
	err := table.Get("ID", id).One(&b)

	if err != nil {
		return *b, err
	}

	return *b, nil
}

func (m *Backlog) Update(id string) (Backlog, error) {
	table = model.DB.Table(TABLE_NAME)
	t := table.Update("ID", id).
		Set("UpdateDate", time.Now())
	if m.GetOwner() != "" {
		t.Set("Owner", m.GetOwner())
	}
	t.Set("Connections", m.GetConnections())

	pRes := new(Backlog)
	err := t.Value(&pRes)

	if err != nil {
		return *pRes, err
	}

	return *pRes, nil
}

func (m *Backlog) Delete(id string) (Backlog, error) {
	oldValue := new(Backlog)

	table = model.DB.Table(TABLE_NAME)
	err := table.Delete("ID", id).OldValue(&oldValue)

	if err != nil {
		return *oldValue, err
	}

	return *oldValue, nil
}

func scan() []Backlog {
	var backlog []Backlog = []Backlog{}
	scanOut, err := model.DB.Client().Scan(&dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		fmt.Println(err.Error())
		return backlog
	}

	for _, scanedBacklog := range scanOut.Items {
		var backlogTmp Backlog

		_ = dynamodbattribute.UnmarshalMap(scanedBacklog, &backlogTmp)
		backlog = append(backlog, backlogTmp)

	}

	return backlog
}

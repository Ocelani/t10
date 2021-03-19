package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"

	"github.com/Ocelani/t10/pkg/auto"
)

func Test_API(t *testing.T) {

	for i := 0; i < 10; i++ {
		test := fmt.Sprintf("teste%v", i*i)

		t.Run(test, func(t *testing.T) {
			token := Auth_add(t, test)
			Auth_authenticate(t, test, token)

			debit := AutoDebit_add(t, test, token)
			AutoDebit_getOne(t, debit.ID.Hex(), token)
			AutoDebit_getOne(t, debit.Name, token)
			AutoDebit_getAll(t, token)

			switch i % 2 {
			case 0:
				AutoDebit_updateStatus(t, debit.ID.Hex(), "approve", token)
				AutoDebit_queryStatus(t, "approved", token)
			case 1:
				AutoDebit_updateStatus(t, debit.ID.Hex(), "reject", token)
				AutoDebit_queryStatus(t, "rejected", token)
			case 2:
				AutoDebit_queryStatus(t, "pending", token)
			}

			AutoDebit_delete(t, debit.ID.Hex(), token)
		})
	}
}

func Auth_add(t *testing.T, test string) string {
	cmd := exec.Command("curl", "-X", "GET", "localhost:80/auth/add/"+test)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("Auth_add")
		t.FailNow()
	}
	token := fmt.Sprintf("X-Auth: %s", cmdOutput.String())

	return token
}

func Auth_authenticate(t *testing.T, test, token string) {
	cmd := exec.Command("curl", "-X", "GET", "-H", token, "localhost:80/auth")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		log.Println(err)
		t.FailNow()
	}
	res := cmdOutput.String()
	if res == "Unauthorized" || res != "Authenticated" {
		t.Log("Auth_authenticate")
		t.FailNow()
	}
}

func AutoDebit_add(t *testing.T, test, token string) auto.Entity {
	cmd := exec.Command(
		"curl", "-X", "POST",
		"-H", token,
		"-d", "name="+test,
		"-d", "amount="+test,
		"-d", "frequency="+test,
		"-d", "status=pending",
		"localhost:80/auto-debit/add",
	)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("AutoDebit_getOne", err)
		t.FailNow()
	}
	out := cmdOutput.String()
	if out == "" {
		t.Log("AutoDebit_getOne", out)
		t.FailNow()
	}
	var res auto.Entity
	if err := json.Unmarshal(cmdOutput.Bytes(), &res); err != nil {
		t.Log("AutoDebit_add JSON: ", err)
	}

	return res
}

func AutoDebit_getOne(t *testing.T, test, token string) {
	cmd := exec.Command(
		"curl", "-X", "GET",
		"-H", token,
		"localhost:80/auto-debit/"+test,
	)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("AutoDebit_getOne", err)
		t.FailNow()
	}
	out := cmdOutput.String()
	if out == "" {
		t.Log("AutoDebit_getOne", out)
		t.FailNow()
	}
	var res auto.Entity
	if err := json.Unmarshal(cmdOutput.Bytes(), &res); err != nil {
		t.Log("AutoDebit_getOne JSON: ", err)
	}
}

func AutoDebit_getAll(t *testing.T, token string) {
	cmd := exec.Command(
		"curl", "-X", "GET",
		"-H", token,
		"localhost:80/auto-debit/all",
	)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("AutoDebit_getAll", err)
		t.FailNow()
	}
	out := cmdOutput.String()
	if out == "" {
		t.Log("AutoDebit_getAll", out)
		t.FailNow()
	}
	var res []auto.Entity
	if err := json.Unmarshal(cmdOutput.Bytes(), &res); err != nil {
		t.Log("AutoDebit_getAll JSON: ", err)
	}
	if len(res) <= 0 {
		t.FailNow()
	}
}

func AutoDebit_updateStatus(t *testing.T, id, status, token string) {
	cmd := exec.Command(
		"curl", "-X", "PUT",
		"-H", token,
		"localhost:80/auto-debit/"+id+"/"+status,
	)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("AutoDebit_updateStatus", err)
		t.FailNow()
	}
	out := cmdOutput.String()
	if out == "" {
		t.Log("AutoDebit_updateStatus", out)
		t.FailNow()
	}
	var res auto.Entity
	if err := json.Unmarshal(cmdOutput.Bytes(), &res); err != nil {
		t.Log("AutoDebit_updateStatus JSON: ", status, err)
	}
	if !strings.Contains(res.Status, status) {
		t.Log("AutoDebit_updateStatus JSON: ", status, res.Status, res)
		t.FailNow()
	}
}

func AutoDebit_queryStatus(t *testing.T, status, token string) {
	cmd := exec.Command(
		"curl", "-X", "GET",
		"-H", token,
		"localhost:80/auto-debit/?status="+status,
	)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("AutoDebit_queryStatus", err)
		t.FailNow()
	}
	out := cmdOutput.String()
	if out == "" {
		t.Log("AutoDebit_queryStatus", out)
		t.FailNow()
	}
	var res []auto.Entity
	if err := json.Unmarshal(cmdOutput.Bytes(), &res); err != nil {
		t.Log("AutoDebit_queryStatus JSON: ", err)
	}
	for _, r := range res {
		if r.Status != status {
			t.Log("AutoDebit_queryStatus")
			t.FailNow()
		}
	}
}

func AutoDebit_delete(t *testing.T, id, token string) {
	cmd := exec.Command(
		"curl", "-X", "DELETE",
		"-H", token,
		"localhost:80/auto-debit/"+id,
	)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		t.Log("AutoDebit_delete", err)
		t.FailNow()
	}
	out := cmdOutput.String()
	if out != id {
		t.Log("AutoDebit_delete", out)
		t.FailNow()
	}
}

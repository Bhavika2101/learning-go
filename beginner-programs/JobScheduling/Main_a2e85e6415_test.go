package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	cron "github.com/robfig/cron/v3"
)

func TestMain_a2e85e6415(t *testing.T) {
	type args struct {
		cronExpr string
		message  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Every hour on the half hour",
			args: args{
				cronExpr: "30 * * * *",
				message:  "Every hour on the half hour",
			},
			want: "Every hour on the half hour\n",
		},
		{
			name: "Runs at 04:30 New York time every day",
			args: args{
				cronExpr: "30 4 * * *",
				message:  "Runs at 04:30 New York time every day",
			},
			want: "Runs at 04:30 New York time every day\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			location, _ := time.LoadLocation("America/New_York")
			c := cron.New(cron.WithLocation(location))
			c.AddFunc(tt.args.cronExpr, func() { fmt.Fprintln(&buf, tt.args.message) })
			c.Start()

			time.Sleep(1 * time.Second)

			got := buf.String()
			if got != tt.want {
				t.Errorf("main() = %v, want %v", got, tt.want)
			}

			c.Stop()
		})
	}
}

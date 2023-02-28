package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ZoomWebhook struct {
	Payload struct {
		AccountID string `json:"account_id"`
		Object    struct {
			UUID        string `json:"uuid"`
			Participant struct {
				LeaveTime         time.Time `json:"leave_time"`
				UserID            string    `json:"user_id"`
				UserName          string    `json:"user_name"`
				RegistrantID      string    `json:"registrant_id"`
				ParticipantUserID string    `json:"participant_user_id"`
				ID                string    `json:"id"`
				LeaveReason       string    `json:"leave_reason"`
				Email             string    `json:"email"`
				ParticipantUUID   string    `json:"participant_uuid"`
			} `json:"participant"`
			ID        string    `json:"id"`
			Type      int       `json:"type"`
			Topic     string    `json:"topic"`
			HostID    string    `json:"host_id"`
			Duration  int       `json:"duration"`
			StartTime time.Time `json:"start_time"`
			Timezone  string    `json:"timezone"`
		} `json:"object"`
	} `json:"payload"`
	EventTs int64  `json:"event_ts"`
	Event   string `json:"event"`
}

func dostuff(c *gin.Context) {
	c.JSON(http.StatusOK, "Did some stuff")
	fmt.Println("Logging to stdout I think")
	fmt.Println(c.Params)
	fmt.Println(c.Request)

	var jresp ZoomWebhook

	if err := c.BindJSON(&jresp); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println("There is an err", err)
	}
	fmt.Println(jresp.Event) // meeting.participant_left
	fmt.Println(jresp.Payload.Object.Participant.UserName)
	fmt.Println(jresp.EventTs)

}

func main() {

	router := gin.Default()
	router.POST("/zoom", dostuff)
	router.Run("localhost:9999")
}

package main

// TODO add basic auth since it's provided in the header
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ZoomWebhook struct {
	Payload struct {
		PlainToken string `json:"plainToken"`
		AccountID  string `json:"account_id"`
		Object     struct {
			UUID        string `json:"uuid"`
			Participant struct {
				LeaveTime         time.Time `json:"leave_time"`
				JoinTime          time.Time `json:"join_time"`
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

type ChallengeResponse struct {
	PlainToken     string `json:"plainToken"`
	EncryptedToken string `json:"encryptedToken"`
}

func dostuff(c *gin.Context) {

	/*fmt.Println("Logging to stdout I think")
	fmt.Println(c.Params)
	fmt.Println(c.Request) */

	var jresp ZoomWebhook

	if err := c.BindJSON(&jresp); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println("There is an err", err)
	}

	// Trying to do the CRC

	if jresp.Payload.PlainToken != "" {
		var crc ChallengeResponse
		crc.PlainToken = jresp.Payload.PlainToken
		//todo load this from env var
		zoom_secret := "020c3V6JTQKxQDucQnkXeg"
		data := jresp.Payload.PlainToken
		// Create a new HMAC by defining the hash type and the key (as byte array)
		h := hmac.New(sha256.New, []byte(zoom_secret))
		h.Write([]byte(data))
		// Get result and encode as hexadecimal string
		crc.EncryptedToken = hex.EncodeToString(h.Sum(nil))
		fmt.Println(crc)
		c.JSON(http.StatusOK, crc)
		return

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

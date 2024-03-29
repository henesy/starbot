package mux

import (
	"github.com/bwmarrin/discordgo"
	"time"
sc	"strconv"
//	"container/list"
//	"fmt"
)


// Channel for posting reminders -- inits in main
var RemChan chan Reminder

// Stores a reminder to be posted to a given user
type Reminder struct {
	NotifyAfter	time.Time
	Reason		string
	User			discordgo.User
	ChannelID	string
	Session		discordgo.Session
}

// Reminder daemon process that gets started in main and listens on RemChan
func Reminders() {
	// TODO -- Should be a heap and more robust (fix config first)
	
	//Rems := list.New()

	// Handle reminders
	for {
		select {
		case r := <- RemChan:
			// Handle new reminder
			Rems.PushBack(r)
			//write()
			
		default:
			// Check for any due reminders
			if Rems != nil {
				if Rems.Front() != nil {
					for e := Rems.Front(); ; e = e.Next() {
						if e == nil {
							break
						}
						
						r, _ := e.Value.(Reminder)
						if time.Now().After(r.NotifyAfter) {
							// If we have passed the time of desired notification
							r.Session.ChannelMessageSend(r.ChannelID, r.User.Mention() + " -- " + r.Reason)
							Rems.Remove(e)
							//write()
						}
					}
				}
			}
		
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Reminds a user about a given reason after a specified time interval has passed
func (m *Mux) RemindMe(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	resp := ""
	
	// remindme 20m dothing
	if len(ctx.Fields) >= 3 {
		var rem Reminder
		periodLong := ctx.Fields[1]
		reasonLong := ctx.Fields[2 : len(ctx.Fields)]
		rem.User = *dm.Author
		rem.ChannelID = dm.ChannelID
		rem.Session = *ds
		
		reason := ""
		for _, v := range reasonLong {
			reason += v
			reason += " "
		}
		rem.Reason = reason
		
		interval := periodLong[len(periodLong)-1]
		period, err := sc.Atoi(periodLong[:len(periodLong)-1])
		if err != nil {
			resp += "Invalid period value."
			goto SEND
		}
		
		var dur time.Duration
		switch interval {
		case 's':
			dur = time.Second
		case 'm':
			dur = time.Minute
		case 'h':
			dur = time.Hour
		case 'd':
			dur = time.Hour * time.Duration(24)
		default:
			resp += "Invalid interval specifier."
			goto SEND
		}

		rem.NotifyAfter = time.Now().Add(time.Duration(period) * dur)
		
		RemChan <- rem
		
		resp += "Ok."
		
	} else {
		resp += "Please specify a time operator in the form [int][type] and a description (20h do thing)."
	}
	
	SEND:
	resp += "\n"

	ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}

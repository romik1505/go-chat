package mapper

// type ChannelData struct {
// 	ID    string  `json:"id"`
// 	Name  string  `json:"name"`
// 	Users UserMap `json:"users"`
// }

const (
	ReceiverTypePerson = ReceiverType("person")
	ReceiverTypeGroup  = ReceiverType("group")
)

type ReceiverType string

type GroupMessage struct {
	Sender  User   `json:"sender"`
	GroupID string `json:"group_id"`
	Text    string `json:"text"`
	Date    string `json:"date"`
}

// TODO: DELETE
type EventMessage struct {
	Type string                `json:"type"`
	Data EventData_SendMessage `json:"data"`
}

// var PersonalMessagesDB = []Message{
// 	{"personal_message", EventData_SendMessage{UserData{"1", "username_1", imgs[0]}, "2", "hello", time.Now()}},
// 	{"personal_message", EventData_SendMessage{UserData{"2", "username_2", imgs[1]}, "1", "hi guy", time.Now()}},
// }

// var GroupMessagesDB = []Message{
// 	{"group_message", EventData_SendMessage{UserData{"1", "username_1", imgs[0]}, "1", "group message 1", time.Now()}},
// 	{"group_message", EventData_SendMessage{UserData{"2", "username_2", imgs[1]}, "2", "group message 2", time.Now()}},
// }

// var imgs = []string{
// 	"https://download-cs.net/steam/avatars/3426.jpg",
// 	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR64s1p-4JgPRXmhXp-81IYUD2d5lrf-gZ3sw&usqp=CAU",
// 	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSmq2s3vVeXJ-NPhm493BCIt3ISGCzd76mNRg&usqp=CAU",
// 	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQgcPn7xBe5xb-JdJJcQQ-aqm-HJ_FnR-zuuQ&usqp=CAU",
// }

// var ClientsDB = UserMap{
// 	"1": {
// 		ID:       "1",
// 		Username: "username_1",
// 		ImgUrl:   imgs[0],
// 	},
// 	"2": {
// 		ID:       "2",
// 		Username: "username_2",
// 		ImgUrl:   imgs[1],
// 	},
// 	"3": {
// 		ID:       "3",
// 		Username: "username_3",
// 		ImgUrl:   imgs[2],
// 	},
// 	"4": {
// 		ID:       "4",
// 		Username: "username_4",
// 		ImgUrl:   imgs[3],
// 	},
// }

// var ChannelsDB = map[string]ChannelData{
// 	"1": {
// 		ID:   "1",
// 		Name: "general",
// 		Users: UserMap{
// 			"3": {
// 				ID:       "3",
// 				Username: "username_3",
// 				ImgUrl:   imgs[2],
// 			},
// 			"4": {
// 				ID:       "4",
// 				Username: "username_4",
// 				ImgUrl:   imgs[3],
// 			},
// 		},
// 	},
// 	"2": {
// 		ID:   "2",
// 		Name: "support",
// 		Users: UserMap{
// 			"1": {
// 				ID:       "1",
// 				Username: "username_1",
// 				ImgUrl:   imgs[0],
// 			},
// 			"2": {
// 				ID:       "2",
// 				Username: "username_2",
// 				ImgUrl:   imgs[1],
// 			},
// 		},
// 	},
// }

// // TODO DELETE STRUCT

// type UserData struct {
// 	ID       string `json:"id,omitempty"`
// 	Username string `json:"username,omitempty"`
// 	ImgUrl   string `json:"img,omitempty"`
// }

package chat

import "sync"

var chatLog []string
var lastMessages []string = make([]string, 10)
var mutex sync.RWMutex

func Add(message string) {
	go rebuildChat(message)
}

func LastMessages() []string {
	// safe async copy
	mutex.RLock()
	copyMessages := make([]string, len(lastMessages))
	copy(copyMessages, lastMessages)
	mutex.RUnlock()

	// Return a copy
	return copyMessages
}

func rebuildChat(message string) {
	// Add last message to queue
	chatLog = append(chatLog, message)

	mutex.Lock()
	lastMessages = lastMessages[1:]
	lastMessages = append(lastMessages, message)
	mutex.Unlock()
}

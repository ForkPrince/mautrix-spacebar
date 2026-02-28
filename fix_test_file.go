package main
import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)
func main() {
    s, _ := discordgo.New("my_token")
    s.Identify.Capabilities = 0
    s.Identify.ClientState = nil
    fmt.Printf("IsUser: %v, ClientState: %v\n", s.IsUser, s.Identify.ClientState)
}

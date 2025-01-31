package states

import (
	"context"
	"fmt"
	"log"

	"math/rand/v2"
	"server/internal/server"
	"server/internal/server/objects"
	"server/pkg/packets"
	"time"
)

type InGame struct {
	client                 server.ClientInterfacer
	player                 *objects.Player
	logger                 *log.Logger
	cancelPlayerUpdateLoop context.CancelFunc
}

func (g *InGame) Name() string {
	return "InGame"
}

func (g *InGame) SetClient(client server.ClientInterfacer) {
	g.client = client
	loggingPrefix := fmt.Sprintf("Client %d [%s]: ", client.Id(), g.Name())
	g.logger = log.New(log.Writer(), loggingPrefix, log.LstdFlags)
}

func (g *InGame) OnEnter() {
	log.Printf("Adding player %s to the shared collection", g.player.Name)
	go g.client.SharedGameObjects().Players.Add(g.player, g.client.Id())

	// Set the initial properties of the player
	g.player.X = rand.Float64() * 1000
	g.player.Y = rand.Float64() * 1000

	// Send the player's initial state to the client
	g.client.SocketSend(packets.NewPlayer(g.client.Id(), g.player))

}

func (g *InGame) HandleMessage(senderId uint64, message packets.Msg) {
	switch message := message.(type) {
	case *packets.Packet_Player:
		g.handlePlayer(senderId, message)
	case *packets.Packet_PlayerInput:
		g.handlePlayerInput(senderId, message)
	case *packets.Packet_Chat:
		g.handleChat(senderId, message)
	}
}

func (g *InGame) OnExit() {
	if g.cancelPlayerUpdateLoop != nil {
		g.cancelPlayerUpdateLoop()
	}
	g.client.SharedGameObjects().Players.Remove(g.client.Id())
}

func (g *InGame) handlePlayer(senderId uint64, message *packets.Packet_Player) {
	if senderId == g.client.Id() {
		g.logger.Println("Received player message from our own client, ignoring")
		return
	}

	g.client.SocketSendAs(message, senderId)
}

func (g *InGame) handlePlayerInput(senderId uint64, message *packets.Packet_PlayerInput) {
	if senderId != g.client.Id() {
		g.logger.Println("Received player direction message from a different client, ignoring")
		return
	}

	g.player.DirectionX = float64(message.PlayerInput.Dx)
	g.player.DirectionY = float64(message.PlayerInput.Dy)
	g.player.Speed = 300

	// If this is the first time receiving a player direction message from our client, start the player update loop
	if g.cancelPlayerUpdateLoop == nil {
		ctx, cancel := context.WithCancel(context.Background())
		g.cancelPlayerUpdateLoop = cancel
		go g.playerUpdateLoop(ctx)
	}
}

func (g *InGame) handleChat(senderId uint64, message *packets.Packet_Chat) {
	if senderId == g.client.Id() {
		g.client.Broadcast(message)
	} else {
		g.client.SocketSendAs(message, senderId)
	}
}

func (g *InGame) playerUpdateLoop(ctx context.Context) {
	const delta float64 = 0.05
	ticker := time.NewTicker(time.Duration(delta*1000) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			g.syncPlayer(delta)
		case <-ctx.Done():
			return
		}
	}
}

func (g *InGame) syncPlayer(delta float64) {
	if g.player.DirectionX == 0 && g.player.DirectionY == 0 {
		return
	}
	newX := g.player.X + g.player.DirectionX*g.player.Speed*delta
	newY := g.player.Y + g.player.DirectionY*g.player.Speed*delta

	g.player.X = newX
	g.player.Y = newY

	updatePlayer := packets.NewPlayer(g.client.Id(), g.player)
	g.client.Broadcast(updatePlayer)
	go g.client.SocketSend(updatePlayer)
}

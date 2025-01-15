extends Node

const packets := preload("res://packets.gd")
const Actor := preload("res://objects/actor/actor.gd")

var _players: Dictionary[int, Actor] = {}

@onready var _line_edit: LineEdit = $UI/LineEdit
@onready var _log: Log = $UI/Log
@onready var _world: Node2D = $World

# Variable to track packet sending status
var is_sending_packet = false

func _ready() -> void:
	WS.connection_closed.connect(_on_ws_connection_closed)
	WS.packet_received.connect(_on_ws_packet_received)
	
	_line_edit.text_submitted.connect(_on_line_edit_text_submitted)

func _handle_chat_msg(sender_id: int, chat_msg: packets.ChatMessage) -> void:
	if sender_id in _players:
		var actor := _players[sender_id]
		_log.chat(actor.actor_name, chat_msg.get_msg())

func _on_line_edit_text_submitted(new_text: String) -> void:
	if is_sending_packet:
		_log.error("Already sending a packet, please wait.")
		return
	
	is_sending_packet = true
	var packet := packets.Packet.new()
	var chat_msg := packet.new_chat()
	chat_msg.set_msg(new_text)
	
	var err := WS.send(packet)
	if err:
		_log.error("Error sending chat message")
	else:
		_log.chat("You", new_text)
	_line_edit.clear()
	is_sending_packet = false
	
func _on_ws_connection_closed() -> void:
	_log.warning("Connection closed")

func _on_ws_packet_received(packet: packets.Packet) -> void:
	var sender_id := packet.get_sender_id()
	
	if packet.has_chat():
		_handle_chat_msg(sender_id, packet.get_chat())
	elif packet.has_player():
		_handle_player_msg(sender_id, packet.get_player())

func _handle_player_msg(sender_id: int, player_msg: packets.PlayerMessage) -> void:
	var actor_id := player_msg.get_id()
	var actor_name := player_msg.get_name()
	var x := player_msg.get_x()
	var y := player_msg.get_y()
	
	var is_player := actor_id == GameManager.client_id
	
	if actor_id not in _players:
		# This is a new player, so we need to create a new actor
		var actor := Actor.instantiate(actor_id, actor_name, x, y, is_player)
		_world.add_child(actor)
		_players[actor_id] = actor
	else:
		# This is an existing player, so we need to update their position
		var actor := _players[actor_id]
		actor.position.x = x
		actor.position.y = y
		
func _update_player(actor_id: int, actor_name: String, x: float, y: float, is_player: bool, input: Vector2) -> void:
	if is_sending_packet:
		return
	
	is_sending_packet = true

	# This is an existing player, so we need to update their position
	var actor := _players[actor_id]
	var packet := packets.Packet.new()
	var player_input_message := packet.new_player_input()
	
<<<<<<< HEAD
<<<<<<< HEAD
	input.x = float(Input.is_action_pressed("ui_right")) - float(Input.is_action_pressed("ui_left"))
	input.y = float(Input.is_action_pressed("ui_down")) - float(Input.is_action_pressed("ui_up"))
	
	# Normalize the input vector
	if input.length() != 0:
		input = input.normalized()
		
=======
	input.x = int(Input.is_action_pressed("ui_right")) - int(Input.is_action_pressed("ui_left"))
	input.y = int(Input.is_action_pressed("ui_down")) - int(Input.is_action_pressed("ui_up"))

	# Normalize the input vector
	if input.length() != 0:
		input = input.normalized()
>>>>>>> parent of 5e94833 (PrototypeV8)
=======
	input.x = int(Input.is_action_pressed("ui_right")) - int(Input.is_action_pressed("ui_left"))
	input.y = int(Input.is_action_pressed("ui_down")) - int(Input.is_action_pressed("ui_up"))

	# Normalize the input vector
	if input.length() != 0:
		input = input.normalized()
>>>>>>> parent of 5e94833 (PrototypeV8)

	player_input_message.set_dx(input.x)
	player_input_message.set_dy(input.y)
	WS.send(packet)
	
<<<<<<< HEAD
<<<<<<< HEAD
	
=======
>>>>>>> parent of 5e94833 (PrototypeV8)
=======
>>>>>>> parent of 5e94833 (PrototypeV8)
	if actor.position.distance_squared_to(Vector2(x, y)) > 100:
		actor.position.x = x
		actor.position.y = y
	
<<<<<<< HEAD
<<<<<<< HEAD
	
=======
>>>>>>> parent of 5e94833 (PrototypeV8)
=======
>>>>>>> parent of 5e94833 (PrototypeV8)
	if not is_player:
		actor.max_speed
	
	is_sending_packet = false

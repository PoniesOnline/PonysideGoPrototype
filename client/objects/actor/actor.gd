extends CharacterBody2D

const packets := preload("res://packets.gd")

const Scene := preload("res://objects/actor/actor.tscn")
const Actor := preload("res://objects/actor/actor.gd")

const max_speed = 300

var input = Vector2.ZERO

@onready var _nameplate: Label = $Nameplate
@onready var _collision_shape: CircleShape2D = $CollisionShape2D.shape
@onready var _camera: Camera2D = $Camera2D
@onready var _sprite: Sprite2D = $Sprite

var actor_id: int
var actor_name: String
var start_x: float
var start_y: float
var is_player: bool

static func instantiate(actor_id: int, actor_name: String, x: float, y: float, is_player: bool) -> Actor:
	var actor := Scene.instantiate()
	actor.actor_id = actor_id
	actor.actor_name = actor_name
	actor.start_x = x
	actor.start_y = y
	actor.is_player = is_player
	
	return actor
	
func _ready():
	position.x = start_x
	position.y = start_y
	
	_nameplate.text = actor_name

func _physics_process(delta):
	if is_player:
		get_input()
		player_movement(delta)
		send_movement()

func get_input():
	input.x = int(Input.is_action_pressed("ui_right")) - int(Input.is_action_pressed("ui_left"))
	input.y = int(Input.is_action_pressed("ui_down")) - int(Input.is_action_pressed("ui_up"))
	return input.normalized()
	
func player_movement(delta):
	velocity = max_speed * input
	move_and_slide()

func send_movement():
	var packet = packets.Packet.new()
	var player_input_message = packet.new_player_input()
	player_input_message.set_dx(int(position.x))
	player_input_message.set_dy(int(position.y))
	WS.send(packet)

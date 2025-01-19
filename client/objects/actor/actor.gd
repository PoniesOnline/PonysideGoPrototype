extends CharacterBody2D

const packets := preload("res://packets.gd")

const Scene := preload("res://objects/actor/actor.tscn")
const Actor := preload("res://objects/actor/actor.gd")

const max_speed = 300

var input = Vector2.ZERO
var last_sent_input = Vector2.ZERO
var previous_x = 0
var previous_y = 0
#var previous_input = Vector2.ZERO #DEBUG
#var previous_velocity = Vector2.ZERO #DEBUG
#var previous_position = Vector2.ZERO #DEBUG
#const POSITION_THRESHOLD = 10 #DEBUG

@onready var _nameplate: Label = $Nameplate
@onready var _collision_shape: CircleShape2D = $CollisionShape2D.shape
@onready var _camera: Camera2D = $Camera2D
@onready var _sprite: Sprite2D = $Sprite
@onready var _animation: AnimationPlayer = $AnimationPlayer

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

func  _physics_process(delta):
	if is_player:
		get_input()
		player_movement(delta)
		if input != last_sent_input:
			send_movement()
			last_sent_input = input
	
	if input.x != previous_x and input.x != 0:
		_sprite.flip_h = input.x == -1
		previous_x = input.x

	if input.x or input.y != 0:
		if not _animation.is_playing():
			_animation.play("walk")
	else:
		if _animation.is_playing():
			_animation.stop()

func get_input():
	input.x = int(Input.is_action_pressed("ui_right")) - int (Input.is_action_pressed("ui_left"))
	input.y = int(Input.is_action_pressed("ui_down")) - int(Input.is_action_pressed("ui_up"))
	input = input#.normalized()
	#Debug Print
	#if input != previous_input:
		#print("Input: ", input)  # Print only when the input changes
		#previous_input = input  # Update the previous input

func player_movement(delta):
	velocity = max_speed * input
	
	#DEBUG PRINT
	#if velocity != previous_velocity:
		#print("Velocity: ", velocity)  # Print only when the velocity changes
		#previous_velocity = velocity  # Update the previous velocity
	
	move_and_slide()
	 # DEBUG PRINT
	#if position.distance_to(previous_position) > POSITION_THRESHOLD:
			#print("Position: ", position)  # Print only when the position changes significantly
			#previous_position = position  # Update the previous position
	

func send_movement():
	var packet = packets.Packet.new()
	var player_input_message = packet.new_player_input()
	player_input_message.set_dx(input.x)
	player_input_message.set_dy(input.y)
	WS.send(packet)

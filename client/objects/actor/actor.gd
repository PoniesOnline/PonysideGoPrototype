extends CharacterBody2D

const packets := preload("res://packets.gd")

const Scene := preload("res://objects/actor/actor.tscn")
const Actor := preload("res://objects/actor/actor.gd")

const max_speed = 300
const accel = 1500
const friction = 2000

var actor_id: int
var actor_name: String
var start_x: float
var start_y: float
var speed: float
var is_player: bool

var input = Vector2.ZERO

@onready var _nameplate: Label = $Nameplate
@onready var _collision_shape: CircleShape2D = $CollisionShape2D.shape
@onready var _camera: Camera2D = $Camera2D
@onready var _sprite: Sprite2D = $Sprite

func _physics_process(delta):
	player_movement(delta)

func get_input():
	input.x = int(Input.is_action_pressed("ui_right"))-int(Input.is_action_pressed("ui_left"))
	input.y = int(Input.is_action_pressed("ui_down"))-int(Input.is_action_pressed("ui_up"))
	return input.normalized()
	
func player_movement(delta):
	input = get_input()
	
	if input == Vector2.ZERO:
		if velocity.length() > (friction*delta):
			velocity -= velocity.normalized() * (friction * delta)
		else:
			velocity = Vector2.ZERO
	else:
		velocity += (input*accel*delta)
		velocity = velocity.limit_length(max_speed)
		
	move_and_slide()

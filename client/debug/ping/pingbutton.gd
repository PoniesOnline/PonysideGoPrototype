extends Button

# Exported variable to set the server URL from the editor
@export var server_url: String = "http://your-server-url/ping"

# Reference to the HTTPRequest node
@onready var http_request: HTTPRequest = $HttpRequest

func _ready() -> void:
	# Connect the button's pressed signal to the _on_button_pressed function
	connect("pressed", self, "_on_button_pressed")
	# Connect the HTTPRequest's request_completed signal to the _on_request_completed function
	http_request.request_completed.connect(_on_request_completed)

func _on_button_pressed() -> void:
	# Get the current time in milliseconds
	var start_time = OS.get_ticks_msec()
	# Send a request to the server URL
	var err = http_request.request(server_url)
	if err != OK:
		print("Ping request failed: ", err)
	else:
		# Store the start time to calculate latency later
		http_request.set_meta("start_time", start_time)

func _on_request_completed(result: int, response_code: int, headers: Array, body: PackedByteArray) -> void:
	# Get the current time in milliseconds
	var end_time = OS.get_ticks_msec()
	# Retrieve the start time stored earlier
	var start_time = http_request.get_meta("start_time")
	# Calculate the latency
	var latency = end_time - start_time

	if response_code == 200:
		print("Ping successful! Latency: ", latency, " ms")
	else:
		print("Ping failed with response code: ", response_code)

<script>
export default {
	props: ['msg']
}
</script>

<template>
	<div class="alert alert-danger" role="alert">
		{{ msg }}
	</div>
</template>

<style scoped>
.alert {
  background-color: #f4f4f4; /* Light gray background */
  border: 1px solid #c0c0c0; /* Light gray border */
  border-radius: 3px;
  padding: 10px 15px;
  font-family: 'Arial', sans-serif; /* Classic font */
  box-shadow: 1px 1px 3px #888; /* Subtle shadow for 3D effect */
  width: 100%; /* Full width for better placement */
  margin: 10px 0; /* Margin for spacing */
  font-size: 14px; /* Classic font size */
}

.alert-danger {
  background-color: #ffdddd; /* Light red background */
  border-color: #d8000c; /* Red border */
  color: #d8000c; /* Red text color */
}

.alert-danger::before {
  content: '!';
  font-weight: bold;
  color: #d8000c; /* Red color for icon */
  margin-right: 10px;
}

.alert .close {
  position: absolute;
  top: 5px;
  right: 10px;
  background-color: #d0d0d0; /* Light gray close button */
  border: 1px solid #a0a0a0; /* Gray border for button */
  border-radius: 3px;
  padding: 0 5px;
  cursor: pointer;
  font-size: 16px;
  color: #000;
}

.alert .close:hover {
  background-color: #b0b0b0; /* Slightly darker gray on hover */
}
</style>

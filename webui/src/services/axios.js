import axios from "axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

function getId(){
    return localStorage["id"]
}

function formatTimestamp(timestamp) {
	const date = new Date(timestamp);
	const options = {
	  year: "numeric",
	  month: "short",
	  day: "2-digit",
	  hour: "2-digit",
	  minute: "2-digit",
	  second: "2-digit",
	}
	return date.toLocaleDateString("en-US", options);
}

function getUsername(){
    return localStorage["username"]
}
const Auth = () => {
	instance.defaults.headers.common['Authorization'] = 'Bearer ' + getId()
}
export {
	Auth,
	formatTimestamp,
	getId,
	getUsername,
	instance as axios
} 

export default instance;
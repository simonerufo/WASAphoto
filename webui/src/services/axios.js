import axios from "axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

function getId(){
    return localStorage["id"]
}

function getUsername(){
    return localStorage["username"]
}
const Auth = () => {
	instance.defaults.headers.common['Authorization'] = 'Bearer ' + getId()
}
export {
	Auth,
	getId,
	getUsername,
	instance as axios
} 

export default instance;
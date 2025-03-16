import axios from "axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

function getId() {
	return localStorage.getItem("id") || null;
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
	};
	return date.toLocaleDateString("en-US", options);
}

function getUsername() {
	return localStorage.getItem("username") || "Guest";
}

instance.interceptors.request.use(
	(config) => {
		const token = getId();
		if (token) {
			config.headers.Authorization = `Bearer ${token}`;
		} else {
			console.warn("Nessun token trovato!");
		}
		console.log("Authorization header:", config.headers.Authorization);
		return config;
	},
	(error) => Promise.reject(error)
);

export { formatTimestamp, getId, getUsername, instance as axios };
export default instance;

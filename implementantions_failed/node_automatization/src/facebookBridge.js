import axios from "axios";
import cookieReader from './cookieReader.js';
import requestHeaders from  './requestHeaders.js';
import JSDOM_pkg from "jsdom"
const {JSDOM} = JSDOM_pkg;

const facebookBridge = {
	login: async(email,pass,uid) => {
		const suffix = "/login/device-based/regular/login/?refsrc=https%3A%2F%2Fmbasic.facebook.com%2F&lwv=100&refid=8";
		//Filling body parameters
		facebookBridge.bodyParameters["email"] = email
		facebookBridge.bodyParameters["pass"] = pass
		
		//Preparing cookies
		const cookies = cookieReader.records[uid];

		//Defining user parameters and URL-encoding them
		const parameters = new URLSearchParams(facebookBridge.bodyParameters);

		//Injecting cookies and including user parameters
		let headers = requestHeaders.POST;
		headers["Content-Length"] = parameters.toString().length; 
		headers["Cookie"] = cookies.map(cookie=>`${cookie.key}=${cookie.value};`).join(" ") 

		//Defining config and making request
		const config = {
			method:"POST",
			baseURL:facebookBridge.prefixURL()+suffix,
			url:suffix,
			headers:headers,
			withCredentials:true,
			validateStatus:(status)=>{
				return status>=200 && status <= 303;
			},
			transformRequest:[(data,headers)=>{
				return data;
			}],
		};
		const instance = axios.create(config);
		const response = await instance.request({data:parameters.toString()})
		console.log(response.data)
	},
	makeReaction:(uid,reaction,url) => {

	},
	get: async(uid,suffix="/") => {
		// Making GET request
		const config = {
			method:"GET",
			url:facebookBridge.prefixURL()+suffix,
			headers:requestHeaders.GET
		}
		const response = await axios(config);
		
		//Getting Cookies
		const cookies = cookieReader.read(response.headers);
		cookieReader.update(uid,cookies);

		//Filling body parameters	
		const jsdom = new JSDOM(response.data);
		const inputs = [...jsdom.window.document.querySelectorAll("input")]
		for(let input of inputs){
			if(Object.keys(facebookBridge.bodyParameters).includes(input.name))
				facebookBridge.bodyParameters[input.name] = input.value;
			console.log(input.name)
		}
	},
	protocol:"https",
	hostname:"mbasic.facebook.com",
	prefixURL: ()=>`${facebookBridge.protocol}://${facebookBridge.hostname}`,
	bodyParameters:{
		"lsd":null,
		"jazoest":null,
		"m_ts":null,
		"li":null,
		"try_number":0,
		"unrecognized_tries":0,
		"email":null,
		"pass":null,
		"login":"Log+In"
	}
}

facebookBridge.get("100044781015699").then(()=>{
	facebookBridge.login("ghzant.y@gmail.com","password(password)","100044781015699");
});
//facebookBridge.makeReaction();

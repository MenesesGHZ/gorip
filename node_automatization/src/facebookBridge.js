import axios from "axios";
import cookieReader from './cookieReader.js';
import requestHeaders from  './requestHeaders.js';
import {promisify} from 'util';
import tough_cookie_pkg from 'tough-cookie';

const {CookieJar} = tough_cookie_pkg;

const facebookBridge = {
	login: async(email,pass,uid) => {
		try{
			const suffix = "login/device-based/regular/login/";

			//Preparing cookies
			const cookieJar = new CookieJar(),
				cookies = cookieReader.records[uid];
			const setCookie = promisify(cookieJar.setCookie.bind(cookieJar));
			for(let cookie of cookies)
				await setCookie(cookie, facebookBridge.prefixURL+suffix);
		
			//Defining user parameters and URL-encoding them
			let parameters = [
				['lsd','AVpJfDk-quE'],
				['email',email],
				['pass',pass]
			];

			//Injecting cookies and including user parameters
			parameters = new URLSearchParams(parameters);
			let headers = requestHeaders.POST;
			headers["Content-Length"] = parameters.toString().length; 
			
			//Defining config and making request
			const config = {
				hostname:facebookBridge.hostname, 
				method:'POST',
				port:443,
				path:`/${suffix}`,
				headers:headers,
				data:parameters,
			};
			/*
			const config = {
				baseURL:facebookBridge.prefixURL,
				headers:headers,
				validateStatus: (status)=>{
					return status>=200 && status <=399;
				},
				maxRedirects:0
			};*/
			
			//Posting
			const instance = await axios.create(config);
			instance.interceptors.response.use(function (response) {
				return response
			}, function (error) {
				console.log({STATUS:error.response.status})
				// Any status codes that falls outside the range of 2xx cause this function to trigger
				// Do something with response error
				return Promise.reject(error);
			});
			instance.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
			instance.post(`/${suffix}`,parameters).then((response)=>{
				console.log(response.request.redirects);
				console.log("\n\n=======================\n\n")
				console.log(response);
			}).catch((error)=>{console.log(error)});




		}catch(error){
			console.log(error)
		}
	},
	makeReaction:(uid,reaction,url) => {

	},
	get: async(uid,suffix="/") => {
		const response = await axios.get(facebookBridge.prefixURL()+suffix);
		const cookie = cookieReader.read(response.headers);
		cookieReader.update(uid,cookie);
		console.log(cookieReader.records);
	},
	protocol:"https",
	hostname:"mbasic.facebook.com",
	prefixURL: ()=>`${facebookBridge.protocol}://${facebookBridge.hostname}`,
	bodyParameters:["lsd","email","pass"]
}
facebookBridge.get("100044781015699").then(()=>{
	//facebookBridge.login("ghzant.y@gmail.com","password(password)","100044781015699");
});
//facebookBridge.makeReaction();

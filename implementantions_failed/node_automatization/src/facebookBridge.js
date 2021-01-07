import axios from "axios";
import cookieReader from './cookieReader.js';
import requestHeaders from  './requestHeaders.js';
import {promisify} from 'util';
import tough_cookie_pkg from 'tough-cookie';

const {CookieJar} = tough_cookie_pkg;


const facebookBridge = {
	login: async(email,pass,uid) => {
		try{
			const suffix = "/login/device-based/regular/login/";

			//Preparing cookies
			const cookieJar = new CookieJar(),
				cookies = cookieReader.records[uid];
			const setCookie = promisify(cookieJar.setCookie.bind(cookieJar));
			for(let cookie of cookies)
				await setCookie(cookie, facebookBridge.prefixURL()+suffix);
		
			//Defining user parameters and URL-encoding them
			const parameters = new URLSearchParams({
				'lsd':'AVpJfDk-quE',
				'email':email,
				'pass':pass
			});

			//Injecting cookies and including user parameters
			let headers = requestHeaders.POST;
			headers["Content-Length"] = parameters.toString().length; 
			
			//Defining config and making request
			axios.defaults.withCredentials = true;
			const config = {
				method:"POST",
				baseURL:facebookBridge.prefixURL(),
				url:suffix,
				headers:headers,
				maxRedirects:0,
				validateStatus:(status)=>{
					console.log("VALIDATING:",status)
					return status>=200 && status <= 303;
				},
				transformRequest:[(data,headers)=>{
					console.log("DATA:",data);
					console.log("HEADERS:",headers);
				}],
				maxBodyLength: 100,
			};
			cookieJar.allowSpecialUseDomain = true;
			cookieJar.enableLooseMode = true;
			const instance = axios.create(config);
			const response = await instance.request({data:parameters.toString()})
			console.log(response.headers)	
			console.log(response.status)	

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
	facebookBridge.login("ghzant.y@gmail.com","password(password)","100044781015699");
});
//facebookBridge.makeReaction();

import requestHeaders from "./requestHeaders.js";

const cookieReader = {
	read: (responseHeaders) => {
		const setCookie = cookieReader._getCookieHeader(responseHeaders,cookieReader.filterArray);
		return setCookie;
	},
	update: (uid,cookie) => {
		if (cookieReader.records[uid]===undefined){
			cookieReader.records[uid] = cookie;
			return;
		}
		for(let key of Object.keys(cookie))
			cookieReader.records[uid][key] = cookie[key];
	},
	_cookieIntoHeader:(headers,cookie) => {
		let cookieString = "";
		for(let keys of Object.keys(cookie))
			cookieString += `${key}=`+`${cookie[key]};`;
		headers["Cookie"] = cookieString;
		return headers;
	},
	_getCookieHeader: (headers,filterArray) => {
		let setCookiePairs = "".concat(...headers["set-cookie"].map(val=>val+";")),
			setCookie={};
		setCookiePairs = setCookiePairs.split(";").map(val=>val.trim().split("=")).slice(0,-1);
		setCookiePairs = setCookiePairs.filter(pair=>filterArray.includes(pair[0]))
		for(let pair of setCookiePairs)
			setCookie[pair[0]] = pair[1];
		return setCookie;
	},
	filterArray:["datr","sb","c_user","xs","fr"],
	records:{}
}

export default cookieReader;

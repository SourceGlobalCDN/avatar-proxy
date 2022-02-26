'use strict';

const Reg = new RegExp("^(/avatar|/gravatar)?/[a-zA-Z0-9]{32}")

exports.main_handler = async (event, context) => {
    const path = event.path;
    if(path.startsWith("/")) {
        return {
            isBase64Encoded: false,
            statusCode: 301,
            headers: {
                "Content-Type":"text/plain",
                "Location": "https://www.sourcegcdn.com"
            },
            body: "301 Permanently Moved"
        }
    }
    if(!Reg.test(path)) {
        return {
            isBase64Encoded: false,
            statusCode: 404,
            headers: {
                "Content-Type":"text/plain"
            },
            body: "file not found, incorrect path"
        }
    }
    if(!path.startsWith("/avatar")) {
        return {
            isBase64Encoded: false,
            statusCode: 302,
            headers: {
                "Content-Type":"text/plain",
                "Location": "/avatar"+path,
                "X-Info": "Please use a url starting with /avatar/"
            },
            body: "302 Moved"
        }
    }
    
};
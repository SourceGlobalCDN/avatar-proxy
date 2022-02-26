"use strict";

const axios = require("axios");

const ImageReg = new RegExp("^(/avatar|/gravatar)/[a-zA-Z0-9]{32}");
const InfoReg = new RegExp("^/[a-zA-Z0-9]{32}\\.(json|xml|php|vcf|qr)");

exports.main_handler = async (event, context) => {
    const path = event.path;
    if (path === "/") {
        return {
            isBase64Encoded: false,
            statusCode: 301,
            headers: {
                "Content-Type": "text/plain",
                Location: "https://www.sourcegcdn.com",
            },
            body: "301 Permanently Moved",
        };
    }
    if (!ImageReg.test(path) && !InfoReg.test(path)) {
        return {
            isBase64Encoded: false,
            statusCode: 404,
            headers: {
                "Content-Type": "text/plain",
            },
            body: "file not found, incorrect path",
        };
    }
    if (path.startsWith("/gravatar")) {
        return {
            isBase64Encoded: false,
            statusCode: 302,
            headers: {
                "Content-Type": "text/plain",
                Location: path.replace("/gravatar/", "/avatar/"),
                "X-Info": "Please use a url starting with /avatar/",
            },
            body: "302 Moved",
        };
    }
    let res;
    try {
        res = await axios("https://www.gravatar.com" + path, {
            method: event.httpMethod.toUpperCase(),
            params: {
                s: event.queryString.hasOwnProperty("s")
                    ? event.queryString.s
                    : undefined,
                d: event.queryString.hasOwnProperty("d")
                    ? event.queryString.d
                    : undefined,
                f: event.queryString.hasOwnProperty("f")
                    ? event.queryString.s
                    : undefined,
                r:
                    event.queryString.hasOwnProperty("s") &&
                    (event.queryString.s === "g" ||
                        event.queryString.s === "pg")
                        ? event.queryString.s
                        : "pg",
            },
            responseType: "arraybuffer",
        });
    } catch (err) {
        console.error("[Axios]", err);
        return {
            isBase64Encoded: false,
            statusCode: 404,
            headers: {
                "Content-Type": "text/plain",
                "X-Info": "Axios error",
            },
            body: "404 Not Found",
        };
    }

    return {
        isBase64Encoded: true,
        statusCode: 200,
        headers: {
            "Content-Type": res.headers["content-type"],
            "X-Info": "Success",
            "use strict";

const axios = require("axios");

const ImageReg = new RegExp("^(/avatar|/gravatar)/[a-zA-Z0-9]{32}");
const InfoReg = new RegExp("^/[a-zA-Z0-9]{32}\\.(json|xml|php|vcf|qr)");

exports.main_handler = async (event, context) => {
    const path = event.path;
    if (path === "/") {
        return {
            isBase64Encoded: false,
            statusCode: 301,
            headers: {
                "Content-Type": "text/plain",
                Location: "https://www.sourcegcdn.com",
            },
            body: "301 Permanently Moved",
        };
    }
    if (!ImageReg.test(path) && !InfoReg.test(path)) {
        return {
            isBase64Encoded: false,
            statusCode: 404,
            headers: {
                "Content-Type": "text/plain",
            },
            body: "file not found, incorrect path",
        };
    }
    if (path.startsWith("/gravatar")) {
        return {
            isBase64Encoded: false,
            statusCode: 302,
            headers: {
                "Content-Type": "text/plain",
                Location: path.replace("/gravatar/", "/avatar/"),
                "X-Info": "Please use a url starting with /avatar/",
            },
            body: "302 Moved",
        };
    }
    let res;
    try {
        res = await axios("https://www.gravatar.com" + path, {
            method: event.httpMethod.toUpperCase(),
            params: {
                s: event.queryString.hasOwnProperty("s")
                    ? event.queryString.s
                    : undefined,
                d: event.queryString.hasOwnProperty("d")
                    ? event.queryString.d
                    : undefined,
                f: event.queryString.hasOwnProperty("f")
                    ? event.queryString.s
                    : undefined,
                r:
                    event.queryString.hasOwnProperty("s") &&
                    (event.queryString.s === "g" ||
                        event.queryString.s === "pg")
                        ? event.queryString.s
                        : "pg",
            },
            responseType: "arraybuffer",
        });
    } catch (err) {
        console.error("[Axios]", err);
        return {
            isBase64Encoded: false,
            statusCode: 404,
            headers: {
                "Content-Type": "text/plain",
                "X-Info": "Axios error",
            },
            body: "404 Not Found",
        };
    }

    return {
        isBase64Encoded: true,
        statusCode: 200,
        headers: {
            "Content-Type": res.headers["content-type"],
            "X-Info": "Success",
            "Link": ImageReg.test(path)?`<https://www.gravatar.com${path}>; rel="canonical"`:undefined
        },
        body: Buffer.from(res.data, "binary").toString("base64"),
    };
};

        },
        body: Buffer.from(res.data, "binary").toString("base64"),
    };
};

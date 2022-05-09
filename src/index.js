"use strict";

const express = require("express");
const app = express();
const port = Number(process.env.PORT) || 9000;

const { GravatarImage } = require("./application/Gravatar/Image");
const { Security } = require("./application/Security");
const { GravatarInfo } = require("./application/Gravatar/Info");
const { GitHubAvatar } = require("./application/GitHub/Avatar");
const { GitHubAvatarByName } = require("./application/GitHub/AvatarByName");

const allowMethod = ["GET", "POST", "HEAD", "OPTIONS"];

app.all("/", (req, res) => {
    res.redirect("https://www.sourcegcdn.com");
    res.end();
});

app.use((res, req, next) => {
    for (let allow of allowMethod) {
        if (res.method.toUpperCase() === allow) {
            next();
            return;
        }
    }
    req.sendStatus(405);
});

app.use(Security);

app.use((req, res, next) => {
    if (
        req.method.toUpperCase() === "GET" &&
        typeof req.referrer === "string" &&
        req.referrer
    ) {
        const url = new URL(req.referrer);
        res.setHeader(
            "Access-Allow-Control-Origin",
            `${url.protocol}://${url.hostname}`
        );
        res.setHeader("Vary", "Origin");
        res.setHeader("Access-Control-Max-Age", 3600);
        res.setHeader("Access-Control-Allow-Credentials", "true");
        res.setHeader("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH");
    }
    res.setHeader("X-Powered-By", "Source Global CDN / www.sourcegcdn.com");
    next();
});

app.all("/ping", (req, res) => {
    res.send({
        code: 0,
        message: "success",
        data: {
            ip: req.ip,
            ua: req.headers["user-agent"],
            time: Date.now(),
        },
    });
});

app.all(new RegExp("^(/avatar|/gravatar)/([a-zA-Z0-9]{32})?$"), GravatarImage);
app.all(new RegExp("^/[a-zA-Z0-9]{32}\\.(json|xml|php|vcf|qr)$"), GravatarInfo);

app.all(new RegExp("^/gh(/[ut])?/\\d+$"), GitHubAvatar);
app.all(new RegExp("^/gh/[a-zA-Z0-9-]+$"), GitHubAvatarByName);

app.listen(port, () => {
    console.log(`Gravatar Proxy listening on port ${port}`);
});

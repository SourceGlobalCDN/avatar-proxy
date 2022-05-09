import express, { NextFunction, Request, Response } from "express";
import GravatarImage from "./application/Gravatar/Image";
import Security from "./application/Security";
import GravatarInfo from "./application/Gravatar/Info";
import GitHubAvatar from "./application/GitHub/Avatar";
import GitHubAvatarByName from "./application/GitHub/AvatarByName";
import { blackList } from "./global";

const app = express();
const port: number =
    typeof process.env.PORT !== "undefined" &&
    !isNaN(parseInt(process.env.PORT))
        ? parseInt(process.env.PORT)
        : 3000;

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

app.use((res, req, next) => {
    for (let allow of allowMethod) {
        if (res.method.toUpperCase() === allow) {
            next();
            return;
        }
    }
    req.sendStatus(405);
});

app.use((req: Request, res: Response, next: NextFunction) => {
    if (
        req.method.toUpperCase() === "GET" &&
        typeof req.headers["referer"] !== "undefined" &&
        req.headers["referer"]
    ) {
        const url = new URL(req.headers["referer"]);
        res.setHeader("Access-Allow-Control-Origin", url.origin);
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
app.all(new RegExp("^/[a-zA-Z\\d]{32}\\.(json|xml|php|vcf|qr)$"), GravatarInfo);

app.all(new RegExp("^/gh(/[ut])?/\\d+$"), GitHubAvatar);
app.all(new RegExp("^/gh/[a-zA-Z\\d-]+$"), GitHubAvatarByName);

app.listen(port, () => {
    console.log(`Gravatar Proxy listening on port ${port}`);
    console.log("Black List:", blackList);
});

import GitHub from "../../middleware/GitHub";
import { Request, Response } from "express";
import { validateBlackList } from "../../global";

const pathReg = new RegExp("^/gh(/.*)");

const GitHubAvatarByName = (req: Request, res: Response) => {
    const path = (req.path.match(pathReg) as string[])[1];
    const userName = (path.match(RegExp("^/([a-zA-Z\\d-]+)")) as string[])[1];
    const apiPath = `/users/${userName}`;

    if (!validateBlackList("githubUser", path.slice(1))) {
        res.sendStatus(403);
        res.end();
        return;
    }

    res.contentType("image/jpeg");

    console.log(
        "[GitHub Avatar by Name]",
        req.method,
        req.originalUrl,
        "API Url:",
        "https://api.github.com" + apiPath
    );

    GitHub.API.get(apiPath, {
        responseType: "json",
    })
        .then((d) => {
            const id = d.data.id;
            console.log(
                "[GitHub Avatar by Name]",
                req.method,
                req.originalUrl,
                "User Id:",
                id
            );

            const size =
                typeof req.query.s !== "undefined" && Number(req.query.s) <= 460
                    ? Number(req.query.s)
                    : 460;

            GitHub.Avatar("/u" + path.replace(userName, id), {
                params: {
                    s: size,
                    v: 4,
                },
                responseType: "arraybuffer",
            })
                .then((r) => {
                    if (r.status === 200) {
                        res.statusCode = 200;
                        res.contentType("image/jpeg");
                        res.setHeader(
                            "Cache-Control",
                            "public, max-age=604800, must-revalidate"
                        );
                        res.send(Buffer.from(r.data, "binary"));
                        res.end();
                    }
                })
                .catch((err) => {
                    console.error("[Axios]", err.message);
                    res.setHeader("Cache-Control", "no-cache");
                    res.sendStatus(404);
                    res.end();
                })
                .then(() => {
                    console.log(
                        "[GitHub Avatar]",
                        req.method,
                        req.originalUrl,
                        res.statusCode
                    );
                });
        })
        .catch((err) => {
            console.error("[Axios]", err.message);
            res.sendStatus(404);
            res.end();
        })
        .then(() => {
            console.log(
                "[GitHub Avatar by Name]",
                req.method,
                req.originalUrl,
                res.statusCode
            );
        });
};

export default GitHubAvatarByName;

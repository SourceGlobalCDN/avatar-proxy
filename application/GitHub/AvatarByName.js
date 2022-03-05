"use strict";

const GitHubAvatarAPI = require("../../middleware/GitHubAPI/Avatar");
const GitHubAPI = require("../../middleware/GitHubAPI/API");

const pathReg = new RegExp("^/gh(/.*)");

/**
 *
 * @param {Request<P, ResBody, ReqBody, ReqQuery, Locals>}req
 * @param {Response<ResBody, Locals>} res
 * @return {void}
 */
exports.GitHubAvatarByName = (req, res) => {
    const path = req.path.match(pathReg)[1];
    const userName = path.match(RegExp("^/([a-zA-Z0-9-]+)"))[1];
    const apiPath = `/users/${userName}`;

    res.contentType("image/jpeg");

    console.log(
        "[GitHub Avatar by Name]",
        req.method,
        req.originalUrl,
        "API Url:",
        "https://api.github.com" + apiPath
    );

    GitHubAPI.get(apiPath, {
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

            GitHubAvatarAPI("/u" + path.replace(userName, id), {
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

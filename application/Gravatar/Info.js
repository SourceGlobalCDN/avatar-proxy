"use strict";

const GravatarAPI = require("../../middleware/GravatarAPI/index");

/**
 *
 * @param {Request<P, ResBody, ReqBody, ReqQuery, Locals>}req
 * @param {Response<ResBody, Locals>} res
 * @return {void}
 */
exports.GravatarInfo = (req, res) => {
    const path = req.path;
    GravatarAPI.post(path, {
        responseType: "arraybuffer",
    })
        .then((r) => {
            if (r.status === 200) {
                res.statusCode = 200;
                res.contentType(r.headers["content-type"]);
                res.send(Buffer.from(r.data, "binary"));
                res.end();
            }
        })
        .catch((err) => {
            console.error("[Axios]", err.message);
            res.sendStatus(404);
            res.end();
        })
        .then(() => {
            console.log(
                "[Gravatar Info]",
                req.method,
                req.originalUrl,
                res.statusCode
            );
        });
};

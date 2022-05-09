import GravatarAPI from "../../middleware/GravatarAPI";
import { Request, Response } from "express";

const GravatarImage = (req: Request, res: Response) => {
    const path = req.path;
    if (path.startsWith("/gravatar")) {
        res.redirect(path.replace("/gravatar/", "/avatar/"));
        res.end();
        return;
    }
    GravatarAPI.get(path, {
        params: {
            s: typeof req.query.s !== "undefined" ? req.query.s : undefined,
            d: typeof req.query.d !== "undefined" ? req.query.d : undefined,
            f: typeof req.query.f !== "undefined" ? req.query.f : undefined,
            r:
                typeof req.query.r !== "undefined" &&
                (req.query.r === "g" || req.query.r === "pg")
                    ? req.query.r
                    : "pg",
        },
        responseType: "arraybuffer",
    })
        .then((r) => {
            if (r.status === 200) {
                res.statusCode = 200;
                res.contentType(r.headers["content-type"]);
                res.setHeader(
                    "Link",
                    `<https://www.gravatar.com${req.originalUrl}>; rel="canonical"`
                );
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
                "[Gravatar Image]",
                req.method,
                req.originalUrl,
                res.statusCode
            );
        });
};

export default GravatarImage;

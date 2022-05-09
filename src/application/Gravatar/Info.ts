import GravatarAPI from "../../middleware/Gravatar";
import { Request, Response } from "express";

const GravatarInfo = (req: Request, res: Response) => {
    const path = req.path;
    GravatarAPI.get(path, {
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

export default GravatarInfo;

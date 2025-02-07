import express, { Router } from "express";
import { dagentLogger } from "@eternalai-dagent/core";
import cors from "cors";

export const getRouter = () => {
    const router = express.Router();

    router.use(cors());
    router.use(express.json());
    router.use(express.urlencoded({ extended: true }));

    return router;
}

export function createApiRouter() {

  const router = getRouter();

  router.get("/hello", (req, res) => {
    res.send("Welcome, this is the REST API!");
  });

  router.get("/api/webhook/twitter-oauth", (req, res) => {
    const code = req.query.code;
    dagentLogger.info(`Twitter OAuth code: `, code);
  });

  return router;
}


export class Direct {
  public app: express.Application;
  private server: any; // Store server instance

  constructor({ routers }: { routers?: Router[] }) {
    dagentLogger.log("Direct constructor");
    this.app = express();
    this.app.use(cors());
    this.app.use(express.json());
    this.app.use(express.urlencoded({ extended: true }));

    // Add routers
    routers?.forEach((router) => {
      this.app.use(router);
    });
  }

  public start(port: number) {
    this.server = this.app.listen(port, () => {
      dagentLogger.success(
          `Server running at http://localhost:${port}`
      );
    });

    const gracefulShutdown = () => {
      dagentLogger.log("Received shutdown signal, closing server...");
      this.server.close(() => {
        dagentLogger.success("Server closed successfully");
        process.exit(0);
      });

      setTimeout(() => {
        dagentLogger.error(
            "Could not close connections in time, forcefully shutting down"
        );
        process.exit(1);
      }, 5000);
    };

    process.on("SIGTERM", gracefulShutdown);
    process.on("SIGINT", gracefulShutdown);
  }

  public stop() {
    if (this.server) {
      this.server.close(() => {
        dagentLogger.success("Server stopped");
      });
    }
  }
}

// const server = new Direct({ routers: [createApiRouter()] });
// server.start(80);

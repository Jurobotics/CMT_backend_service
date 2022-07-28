import { Router } from "oak";

const router = new Router({
    prefix: "/queue",
});

router.get('/', async (ctx) => {
    const sock = await ctx.upgrade();

    sock.onopen = () => {
        sock.send('Hello World!');
    }

    sock.onmessage = (event) => {
        console.log(event.data);
        switch (event.data) {
            case 'ping':
                sock.send('pong');
                break;
            case 'pong':
                sock.send('ping');
                break;
            default:
                sock.send('unknown');
                break;
        }
    }
});

export default router;


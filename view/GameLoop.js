export default class GameLoop {
    constructor({
        renderer,
        updater,
    }) {
        this.duration = 0;
        this.rate = 60.0;
        this.tickInterval = 1000 / this.rate; // 0.016666s = 16.66666ms

        this.lastTick = 0;
        this.renderer = renderer;
        this.updater = updater;

    }

    start() {
        const current = performance.now();
        if (this.lastTick === 0) {
            this.lastTick = current
        }
        this.duration += current - this.lastTick;
        this.lastTick = current;
        if (this.duration < this.tickInterval) {
            requestAnimationFrame(() => {this.start()});
            return;
        }
        while (this.duration > this.tickInterval) {
            this.updater();
            this.duration -= this.tickInterval
        }
        this.renderer();
        requestAnimationFrame(() => {this.start()});
    }

}
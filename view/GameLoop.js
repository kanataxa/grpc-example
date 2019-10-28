export default class GameLoop {
    constructor(renderer) {
        this.duration = 0;
        this.rate = 60.0;
        this.tickInterval = 1000 / this.rate; // 0.016666s = 16.66666ms

        this.lastTick = 0;
        this.renderer = renderer;
        console.log(this)

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
            // game state update()
            this.duration -= this.tickInterval
        }
        this.renderer()
        requestAnimationFrame(() => {this.start()});
    }

}
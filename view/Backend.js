export default class Backend {
    constructor() {
        this.interval = 100; // 100ms
        this._requestInputs = [];
    }
    add(key) {
        this._requestInputs.push(key.keyCode);
    }
    
}
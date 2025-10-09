export namespace interfaces {
	
	export class RecordingMetaData {
	    FileName: string;
	    Started: time.Time;
	
	    static createFrom(source: any = {}) {
	        return new RecordingMetaData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.FileName = source["FileName"];
	        this.Started = this.convertValues(source["Started"], time.Time);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace models {
	
	export class LiveAudioEventData {
	    loudnessPercentage: number;
	
	    static createFrom(source: any = {}) {
	        return new LiveAudioEventData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.loudnessPercentage = source["loudnessPercentage"];
	    }
	}
	export class RecordingInfo {
	    fileName: string;
	    durationSeconds: number;
	
	    static createFrom(source: any = {}) {
	        return new RecordingInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.durationSeconds = source["durationSeconds"];
	    }
	}
	export class RunningEventData {
	    fileName: string;
	    secondsRunning: number;
	
	    static createFrom(source: any = {}) {
	        return new RunningEventData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.secondsRunning = source["secondsRunning"];
	    }
	}

}

export namespace services {
	
	export class Settings {
	    RecordingsDirectory: string;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RecordingsDirectory = source["RecordingsDirectory"];
	    }
	}

}

export namespace time {
	
	export class Time {
	
	
	    static createFrom(source: any = {}) {
	        return new Time(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}


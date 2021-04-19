import {Timestamp} from 'rxjs';

export class Call {
  Id: number;
  // tslint:disable-next-line:variable-name
  Name: string;
  From: string;
  To: string;
  // tslint:disable-next-line:variable-name
  Start_timestamp: Timestamp<string>;
  // tslint:disable-next-line:variable-name
  End_timestamp: Timestamp<string>;
  Duration: any;
}


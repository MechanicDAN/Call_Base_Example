import {Component, OnInit, ViewChild} from '@angular/core';
import {HttpService} from '../../http.service';
import {Call} from '../../structures/call';
import {TableComponent} from '../table/table.component';
import {User} from '../../structures/user';
import {Column} from '../../structures/column';

@Component({
  selector: 'app-calls',
  templateUrl: './calls.component.html',
  styleUrls: ['./calls.component.css']
})
export class CallsComponent implements OnInit {
  @ViewChild(TableComponent)
  private tableComponent: TableComponent;
  columns = [new Column('Name', 'call.th.user', ''), new Column('From', 'call.th.from', ''), new Column('To', 'call.th.to', ''),
    new Column('Start_timestamp', 'call.th.startTimestamp', 'date'), new Column('End_timestamp', 'call.th.endTimestamp', 'date'),
    new Column('Duration', 'call.th.duration', 'time')];
  calls: Call[];
  users: User[];
  currentUser = '';
  sortBy = '';
  sortType = '';

  constructor(private httpService: HttpService) {}

  ngOnInit(): void {
    this.httpService.getUsers().subscribe( (data: User[]) => {
      this.users = data;
    });
    this.httpService.getCalls().subscribe((data: Call[]) => {
      this.calls = data;
      console.log(this.calls); });
  }
  submit(): void{
    this.httpService
      .tableGetCalls(this.currentUser, this.tableComponent.sortBy, this.tableComponent.sortType)
      .subscribe((data: Call[]) => {
        this.calls = data; });
    console.log(this.calls);
  }
  setSort(sort): void{
    this.sortBy = sort[0];
    this.sortType = sort[1];
  }
}

import {Component, Input, OnInit, Output, EventEmitter } from '@angular/core';
import {Column} from '../../structures/column';
import {faSort, faSortDown, faSortUp} from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-table',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.css']
})
export class TableComponent implements OnInit {
  @Output() change = new EventEmitter();
  @Output() setSort = new EventEmitter<string[]>();
  @Output() edit = new EventEmitter<string>();
  @Input() inputColumns: Column[];
  @Input() data: any[];
  @Input() showId: boolean;
  columns: Column[] = [];
  sortBy = '';
  sortType = '';
  faSort = faSort;
  faSortDown = faSortDown;
  faSortUp = faSortUp;
  keepOrder = (a, b) => a;

  constructor() {}

  ngOnInit(): void {
      this.columns = this.inputColumns;
  }

  sort(item: Column): void {
    switch (item.type) {
      case 'ASC': {
        this.resetColumns();
        this.sortBy = item.name;
        this.sortType = 'DESC';
        item.type = 'DESC';
        break;
      }
      case 'DESC': {
        this.resetColumns();
        this.sortBy = item.name;
        this.sortType = '';
        item.type = '';
        break;
      }
      case '': {
        this.resetColumns();
        this.sortBy = item.name;
        this.sortType = 'ASC';
        item.type = 'ASC';
        break;
      }
    }
    this.setSort.emit([this.sortBy, this.sortType]);
    this.change.emit();
  }

  resetColumns(): void{
    for (const item of this.columns){
      item.type = '';
    }
  }

  editEvent(id): void {
    this.edit.emit(id);
  }
}

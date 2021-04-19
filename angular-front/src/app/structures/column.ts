export class Column {
  name: string;
  text: string;
  type: string;
  format: string;

  constructor(name: string, text: string, format: string) {
    this.name = name;
    this.text = text;
    this.format = format;
    this.type = '';
  }
}

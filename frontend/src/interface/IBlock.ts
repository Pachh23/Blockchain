export interface BlockInterface {
  index: number;
  timestamp: string;  // ใช้เป็น string เพื่อให้แสดงเวลาในรูปแบบที่ต้องการ
  previousHash: string;
  data: string;
  hash: string;
}

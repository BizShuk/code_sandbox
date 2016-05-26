基本概念
ref. 排序演算法列表 http://zh.wikipedia.org/wiki/%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95
http://jashliao.pixnet.net/blog/post/154600880-c%2B%2B%E7%B6%93%E5%85%B8%E7%9A%847%E7%A8%AE%E6%8E%92%E5%BA%8F%E6%BC%94%E7%AE%97%E6%B3%95


void ExchangeSort(int* array, int size);O(n2)	穩定
	n輪 從第一個跟其他比較 比較小 馬上交換 		=>交換次數很多

void SelectSort(int* array, int size);	O(n2)	不穩定
	從第一個跟其他比較 一輪比完才交換		=>優化的ExchangeSort

void InsertSort(int* array, int size);	O(n2)	穩定
	跑 n 輪 i: 1 -> array_size , j: i -> 1 , array size: 2 -> n , 兩兩比較交換
	適用於數列較小的情況 最耗時間

void BubbleSort(int* array, int size);	O(n2)	穩定
	跑 n 輪 i: 0 -> n-1 , j: 1 -> array_size , array size: n -> 2, 兩兩比較交換(會將最大or最小放在最後面) 

void MergeSort(int* array, int left, int right);
void Merge(int* array, int left, int mid, int right);
void HeapSort(int *array, int size);
void HeapAjust(int *array, int toAjust, int size);
void QuickSort(int *array, int left, int right);
void TreeSort(int *array, int size);
void TreeMidRetrival(int* array, int* temp, int pos, int* lChild, int* rChild);
void Swap(int* array, int x, int y);

BubbleSort

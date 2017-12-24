package worker

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var content string

func init() {
	content = `|                                                                        -----Statements Bill----                                                                        
	|                                                                      -----My NAME IS CHANGE-----                                                                       
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|Account No：61188803                                                      Statement Date：2017-12-12 to 2017-12-12
	|Account Name：邦吉（上海）谷物三部                                        Bill Date：2017-12-13
	|Account Type：Sub Account                                                 
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|                                                                          Journal Description                                                                           
	|   Date   |     Cash In      |     Cash Out       |   Type   |Currency|Remarks |
	|2017-12-12|              0.00|      1,000,000.00  |  Amount  |  CNY   |        |
	| Summary  |              0.00|      1,000,000.00  |          |  CNY   |        |
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|                                                                           Trade Confirmation                                                                           
	|  Date    | Market |        Product         |    Contract    | Open/Close  | FocusClose  |Buy/Sale|MatchQty| Match Price  |    Premium     |     Fee      |Currency|Remarks |         Time         |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   10   |  1706.0000000|            0.00|         17.00|  CNY   |        | 2017-12-12 09:30:15  |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   10   |  1704.0000000|            0.00|         17.00|  CNY   |        | 2017-12-12 11:06:14  |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   3    |  1704.0000000|            0.00|          5.10|  CNY   |        | 2017-12-12 09:31:43  |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   7    |  1704.0000000|            0.00|         11.90|  CNY   |        | 2017-12-12 09:31:43  |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   1    |  1703.0000000|            0.00|          1.70|  CNY   |        | 2017-12-12 11:29:50  |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   9    |  1703.0000000|            0.00|         15.30|  CNY   |        | 2017-12-12 11:29:55  |
	|2017-12-12|  DCE   |           C            |      1801      |    Close    |             |  Buy   |   10   |  1705.0000000|            0.00|         17.00|  CNY   |        | 2017-12-12 09:31:25  |
	| Summary  |        |                        |                |             |             |        |   50   |              |            0.00|         85.00|  CNY   |        |
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|                                                                            Close Positions                                                                             
	|   Date   | Market |        Product         |    Contract    |Buy/Sale|  Qty   |  Open Price  | Close Price  | Settle Price | Current Profit |Currency|
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   10   |  1694.0000000|  1706.0000000|  1707.0000000|       -1,200.00|  CNY   |
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   10   |  1694.0000000|  1705.0000000|  1707.0000000|       -1,100.00|  CNY   |
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   3    |  1695.0000000|  1704.0000000|  1707.0000000|         -270.00|  CNY   |
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   7    |  1695.0000000|  1704.0000000|  1707.0000000|         -630.00|  CNY   |
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   10   |  1695.0000000|  1704.0000000|  1707.0000000|         -900.00|  CNY   |
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   1    |  1696.0000000|  1703.0000000|  1707.0000000|          -70.00|  CNY   |
	|2017-11-21|  DCE   |           C            |      1801      |  Buy   |   9    |  1696.0000000|  1703.0000000|  1707.0000000|         -630.00|  CNY   |
	| Summary  |        |                        |                |        |   50   |              |              |              |       -4,800.00|  CNY   |
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|                                                                        Detailed Open Positions                                                                         
	|   Date   | Market |        Product         |    Contract    |  Buy   |  Sale  | Match Price  |Last Settlement|Settlement Price|  Current Profit|Option MarketValue|Currency|
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   14   |  2300.0000000|   2353.0000000|    2348.0000000|       -6,720.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   2    |  2300.0000000|   2353.0000000|    2348.0000000|         -960.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   2    |  2300.0000000|   2353.0000000|    2348.0000000|         -960.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   79   |  2300.0000000|   2353.0000000|    2348.0000000|      -37,920.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |  100   |  2299.0000000|   2353.0000000|    2348.0000000|      -49,000.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   8    |  2300.0000000|   2353.0000000|    2348.0000000|       -3,840.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   2    |  2300.0000000|   2353.0000000|    2348.0000000|         -960.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   4    |  2300.0000000|   2353.0000000|    2348.0000000|       -1,920.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   46   |  2300.0000000|   2353.0000000|    2348.0000000|      -22,080.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2300.0000000|   2353.0000000|    2348.0000000|         -480.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   32   |  2300.0000000|   2353.0000000|    2348.0000000|      -15,360.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   10   |  2303.0000000|   2353.0000000|    2348.0000000|       -4,500.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   6    |  2303.0000000|   2353.0000000|    2348.0000000|       -2,700.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   50   |  2303.0000000|   2353.0000000|    2348.0000000|      -22,500.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   1    |  2303.0000000|   2353.0000000|    2348.0000000|         -450.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   5    |  2303.0000000|   2353.0000000|    2348.0000000|       -2,250.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   20   |  2303.0000000|   2353.0000000|    2348.0000000|       -9,000.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   8    |  2303.0000000|   2353.0000000|    2348.0000000|       -3,600.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   59   |  2302.0000000|   2353.0000000|    2348.0000000|      -27,140.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   5    |  2302.0000000|   2353.0000000|    2348.0000000|       -2,300.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   3    |  2302.0000000|   2353.0000000|    2348.0000000|       -1,380.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   30   |  2302.0000000|   2353.0000000|    2348.0000000|      -13,800.00|        0         |  CNY    |
	|2017-11-20|  ZCE   |           RM           |      805       |   0    |   3    |  2302.0000000|   2353.0000000|    2348.0000000|       -1,380.00|        0         |  CNY    |
	|2017-11-21|  DCE   |           C            |      1801      |   0    |   10   |  1696.0000000|   1707.0000000|    1705.0000000|         -900.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1702.0000000|   1707.0000000|    1705.0000000|         -300.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1702.0000000|   1707.0000000|    1705.0000000|         -300.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1702.0000000|   1707.0000000|    1705.0000000|         -300.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1703.0000000|   1707.0000000|    1705.0000000|         -200.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1703.0000000|   1707.0000000|    1705.0000000|         -200.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1704.0000000|   1707.0000000|    1705.0000000|         -100.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1705.0000000|   1707.0000000|    1705.0000000|            0.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1705.0000000|   1707.0000000|    1705.0000000|            0.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1704.0000000|   1707.0000000|    1705.0000000|         -100.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1705.0000000|   1707.0000000|    1705.0000000|            0.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1705.0000000|   1707.0000000|    1705.0000000|            0.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1706.0000000|   1707.0000000|    1705.0000000|          100.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1706.0000000|   1707.0000000|    1705.0000000|          100.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1706.0000000|   1707.0000000|    1705.0000000|          100.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1707.0000000|   1707.0000000|    1705.0000000|          200.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1708.0000000|   1707.0000000|    1705.0000000|          300.00|        0         |  CNY    |
	|2017-11-23|  DCE   |           C            |      1801      |   0    |   10   |  1709.0000000|   1707.0000000|    1705.0000000|          400.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   10   |  1707.0000000|   1707.0000000|    1705.0000000|          200.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   10   |  1708.0000000|   1707.0000000|    1705.0000000|          300.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   10   |  1709.0000000|   1707.0000000|    1705.0000000|          400.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   10   |  1710.0000000|   1707.0000000|    1705.0000000|          500.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   10   |  1710.0000000|   1707.0000000|    1705.0000000|          500.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   5    |  1711.0000000|   1707.0000000|    1705.0000000|          300.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   1    |  1711.0000000|   1707.0000000|    1705.0000000|           60.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   1    |  1711.0000000|   1707.0000000|    1705.0000000|           60.00|        0         |  CNY    |
	|2017-11-24|  DCE   |           C            |      1801      |   0    |   3    |  1711.0000000|   1707.0000000|    1705.0000000|          180.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   10   |  1720.0000000|   1707.0000000|    1705.0000000|        1,500.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   10   |  1720.0000000|   1707.0000000|    1705.0000000|        1,500.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	|2017-11-29|  DCE   |           C            |      1801      |   0    |   1    |  1720.0000000|   1707.0000000|    1705.0000000|          150.00|        0         |  CNY    |
	| Summary  |        |                        |                |   0    |  770   |              |               |                |    -230200     |        0         |  CNY    |
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|                                                                        Gathered Open Positions                                                                         
	|  Market  |        Product         |    Contract    |  Buy   |  Sale  | Match Price  |Settlement Price| Position Profit|Option MarketValue|     Margin     |Currency|
	|   ZCE    |           RM           |      805       |   0    |  500   |  2300.8000000|    2348.0000000|     -236,000.00|              0.00|    1,526,200.00|  CNY   |
	|   DCE    |           C            |      1801      |   0    |  270   |  1707.1500000|    1705.0000000|        5,800.00|              0.00|      322,245.00|  CNY   |
	| Summary  |                        |                |   0    |  770   |              |                |     -230,200.00|              0.00|    1,848,445.00|  CNY   |
	--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
	|                                                                          Financial Situation                                                                           
	|Currency          |      BaseCurrency|               CNY|                  |                  |                  |                  |                  
	|Exchange          |            1.0000|            1.0000|                  |                  |                  |                  |                  
	|Opening           |      4,337,763.00|      4,337,763.00|                  |                  |                  |                  |                  
	|Deposit/Withdrawal|     -1,000,000.00|     -1,000,000.00|                  |                  |                  |                  |                  
	|Journal           |              0.00|              0.00|                  |                  |                  |                  |                  
	|Commissions       |             85.00|             85.00|                  |                  |                  |                  |                  
	|Trading           |         -4,800.00|         -4,800.00|                  |                  |                  |                  |                  
	|Delivery          |              0.00|              0.00|                  |                  |                  |                  |                  
	|Option            |              0.00|              0.00|                  |                  |                  |                  |                  
	|Closing           |      3,332,878.00|      3,332,878.00|                  |                  |                  |                  |                  
	|Unrealized        |              0.00|              0.00|                  |                  |                  |                  |                  
	|Floating          |       -230,200.00|       -230,200.00|                  |                  |                  |                  |                  
	|Equity            |      3,102,678.00|      3,102,678.00|                  |                  |                  |                  |                  
	|PreEquity         |      4,071,063.00|      4,071,063.00|                  |                  |                  |                  |                  
	|Option            |              0.00|              0.00|                  |                  |                  |                  |                  
	|Account           |      3,102,678.00|      3,102,678.00|                  |                  |                  |                  |                  
	|Initial           |      1,848,445.00|      1,848,445.00|                  |                  |                  |                  |                  
	|Maintenance       |      1,848,445.00|      1,848,445.00|                  |                  |                  |                  |                  
	|Excess            |      1,254,233.00|      1,254,233.00|                  |                  |                  |                  |                  
	|Account           |              0.00|              0.00|                  |                  |                  |                  |                  
	----------------
	`

}

func TestProcess(t *testing.T) {
	src := "./"
	destination := "./"
	filename := "_test.txt"
	filepath := src + "/" + filename

	err := ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Println("Create test file error")
		os.Exit(1)
	}
	defer os.Remove(filepath)

	s, _ := process(filename, src, destination)
	defer os.Remove(destination + "/" + s)

	if len(s) <= 0 {
		t.Errorf("Expected csv file created, but not")
	}

	if !strings.Contains(s, "61188803") {
		t.Errorf("Expected csv file name contains '61188803', but not")
	}

}

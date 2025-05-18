package main

import (
   "encoding/csv"
   "fmt"
   "os"
   "sort"
   "strconv"
)

type Product struct {
   ID       int
   Name     string
   Qty      int
   Price    int
   Discount int
}

func (p Product) Cost() int {
   discountFactor := float64(100-p.Discount) / 100
   return int(float64(p.Price*p.Qty) * discountFactor)
}

var products []Product

func init() {
   products = LoadCSV()
}

func LoadCSV() []Product {
   file, err := os.Open("data.csv")
   if err != nil {
      fmt.Println("Error membuka file:", err)
      return []Product{}
   }

   defer file.Close()
   reader := csv.NewReader(file)
   header, err := reader.Read()
   if err != nil {
      fmt.Println("Error membaca header:", err)
      return []Product{}
   }
   fmt.Println("Header:", header)

   records, err := reader.ReadAll()
   if err != nil {
      fmt.Println("Error membaca isi file:", err)
      return []Product{}
   }

   var products []Product
   for _, row := range records {
      if len(row) < 5 {
         fmt.Println("Baris tidak valid:", row)
         	continue
      }
      id, _ := strconv.Atoi(row[0])
      name := row[1]
      price, _ := strconv.Atoi(row[2])
      qty, _ := strconv.Atoi(row[3])
      discount, _ := strconv.Atoi(row[4])

      products = append(products, Product{
         ID:       id,
         Name:     name,
         Price:    price,
         Qty:      qty,
         Discount: discount,
      })
   }
   return products
}

func findMostExpensive(products []Product) Product {
   max := products[0]
   for _, p := range products {
      if p.Cost() > max.Cost() {
         max = p
      }
   }
   return max
}

func findCheapest(products []Product) Product {
   min := products[0]
   for _, p := range products {
      if p.Cost() < min.Cost() {
         min = p
      }
   }
   return min
}

func filterExpensive(products []Product) []Product {
   var result []Product
   for _, p := range products {
      if p.Cost() > 500000 {
         result = append(result, p)
      }
   }
   return result
}

func binarySearchByCost(products []Product, target int) *Product {
   sort.Slice(products, func(i, j int) bool {
      return products[i].Cost() < products[j].Cost()
   })

   low, high := 0, len(products)-1
   for low <= high {
      mid := (low + high) / 2
      midCost := products[mid].Cost()
      if midCost == target {
         return &products[mid]
      }else if midCost < target {
         low = mid + 1
      }else{
         high = mid - 1
      }
   }
   return nil
}

func main() {
	fmt.Println("\n======================= DATA PRODUK ===========================")
	fmt.Printf("ID   | Nama            | Harga       | Qty | Diskon | Biaya\n")
	fmt.Println("-----+-----------------+-------------+-----+--------+---------")

	for _, p := range products {
		fmt.Printf("%-4d | %-15s | %11d | %3d | %6s | %7d\n",
			p.ID, p.Name, p.Price, p.Qty, fmt.Sprintf("%d%%", p.Discount), p.Cost())
	}

	fmt.Println("\n=== PRODUK DENGAN BIAYA TERMAHAL ===")
	pMax := findMostExpensive(products)
	fmt.Printf("%-4d | %-15s | %11d | %3d | %6s | %7d\n",
		pMax.ID, pMax.Name, pMax.Price, pMax.Qty, fmt.Sprintf("%d%%", pMax.Discount), pMax.Cost())

	fmt.Println("\n=== PRODUK DENGAN BIAYA TERMURAH ===")
	pMin := findCheapest(products)
	fmt.Printf("%-4d | %-15s | %11d | %3d | %6s | %7d\n",
		pMin.ID, pMin.Name, pMin.Price, pMin.Qty, fmt.Sprintf("%d%%", pMin.Discount), pMin.Cost())

	fmt.Println("\n=== PRODUK DENGAN BIAYA > 500.000 ===")
	expensive := filterExpensive(products)
	for _, p := range expensive {
		fmt.Printf("%-4d | %-15s | %11d | %3d | %6s | %7d\n",
			p.ID, p.Name, p.Price, p.Qty, fmt.Sprintf("%d%%", p.Discount), p.Cost())
	}

	fmt.Println("\n=== MENCARI PRODUK DENGAN BIAYA 35.000 ===")
	result := binarySearchByCost(products, 35000)
	if result != nil {
		fmt.Printf("Ditemukan: %-4d | %-15s | %11d | %3d | %6s | %7d\n",
			result.ID, result.Name, result.Price, result.Qty, fmt.Sprintf("%d%%", result.Discount), result.Cost())
	}else{
		fmt.Println("Produk tidak ditemukan.")
	}
}
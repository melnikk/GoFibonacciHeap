package fibHeap

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"math/rand"
	"time"
)

var _ = Describe("Tests of fibHeap", func() {
	var (
		heap        *FibHeap
		anotherHeap *FibHeap
	)

	Context("behaviour tests", func() {
		BeforeEach(func() {
			heap = NewFibHeap()
		})

		AfterEach(func() {
			heap = nil
		})

		It("Given an empty fibHeap, when call Minimum api, it should return nil.", func() {
			Expect(heap.Minimum()).Should(BeNil())
		})

		It("Given a empty fibHeap, when call Insert api with a nil value, it should return error.", func() {
			Expect(heap.Insert(nil)).Should(HaveOccurred())
		})

		It("Given a empty fibHeap, when call Insert api with a negetive infinity key, it should return error.", func() {
			demo := new(demoStruct)
			demo.tag = 1000
			demo.key = math.Inf(-1)
			demo.value = fmt.Sprint(1000)

			Expect(heap.Insert(demo)).Should(HaveOccurred())
		})

		It("Given a fibHeap inserted multiple values, when call Minimum api, it should return the minimum value inserted.", func() {
			min := math.Inf(1)
			rand.Seed(time.Now().Unix())
			for i := 0; i < 10000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
				if demo.key < min {
					min = demo.key
				}
			}

			Expect(heap.Num()).Should(BeEquivalentTo(10000))
			Expect(heap.Minimum().(*demoStruct).key).Should(Equal(min))
			Expect(heap.Minimum().(*demoStruct).value).Should(Equal(fmt.Sprint(min)))
			Expect(heap.Num()).Should(BeEquivalentTo(10000))
		})

		It("Given an empty fibHeap, when call ExtractMin api, it should return nil.", func() {
			Expect(heap.ExtractMin()).Should(BeNil())
		})

		It("Given a fibHeap inserted multiple values, when call ExtractMin api, it should extract the minimum value inserted.", func() {
			rand.Seed(time.Now().Unix())
			for i := 0; i < 10000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}

			Expect(heap.Num()).Should(BeEquivalentTo(10000))
			lastKey := heap.Minimum().(*demoStruct).key
			for i := 0; i < 10000; i++ {
				extracted := heap.ExtractMin().(*demoStruct)
				Expect(extracted.key).Should(BeNumerically(">=", lastKey))
				Expect(heap.Num()).Should(BeEquivalentTo(9999 - i))
				lastKey = extracted.key
			}
			Expect(heap.Num()).Should(BeEquivalentTo(0))
		})

		It("Given a fibHeap, when call DecreaseKey api with a nil value, it should return error.", func() {
			Expect(heap.DecreaseKey(nil)).Should(HaveOccurred())
		})

		It("Given a fibHeap inserted multiple values, when call DecreaseKey api with a non-exists value, it should return error.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i)
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}

			decreaseDemo := new(demoStruct)
			decreaseDemo.tag = 1000
			decreaseDemo.key = float64(999)
			decreaseDemo.value = fmt.Sprint(decreaseDemo.key)

			Expect(heap.DecreaseKey(decreaseDemo)).Should(HaveOccurred())
			Expect(heap.Num()).Should(BeEquivalentTo(1000))
		})

		It("Given a fibHeap with a value, when call DecreaseKey api with a negetive infinity key, it should return error.", func() {
			demo := new(demoStruct)
			demo.tag = 1000
			demo.key = float64(1000)
			demo.value = fmt.Sprint(1000)
			heap.Insert(demo)

			demo.key = math.Inf(-1)
			Expect(heap.DecreaseKey(demo)).Should(HaveOccurred())
		})

		It("Given a fibHeap inserted multiple values, when call DecreaseKey api with a greater key, it should return error.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i)
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}

			decreaseDemo := new(demoStruct)
			decreaseDemo.tag = 999
			decreaseDemo.key = float64(1000)
			decreaseDemo.value = fmt.Sprint(decreaseDemo.key)

			Expect(heap.DecreaseKey(decreaseDemo)).Should(HaveOccurred())
			Expect(heap.Num()).Should(BeEquivalentTo(1000))
		})

		It("Given a fibHeap inserted multiple values, when call DecreaseKey api with a smaller key, it should decrease the key of the value in the heap.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i + 1000)
				demo.value = fmt.Sprint(i)
				heap.Insert(demo)
			}
			heap.ExtractMin()
			for i := 999; i >= 1; i-- {
				decreaseDemo := new(demoStruct)
				decreaseDemo.tag = i
				decreaseDemo.key = float64(i)
				decreaseDemo.value = fmt.Sprint(i)
				Expect(heap.DecreaseKey(decreaseDemo)).ShouldNot(HaveOccurred())
			}
			Expect(heap.Num()).Should(BeEquivalentTo(999))

			for i := 1; i < 1000; i++ {
				value := heap.ExtractMin()
				Expect(value.(*demoStruct).tag).Should(BeEquivalentTo(i))
				Expect(value.(*demoStruct).key).Should(BeEquivalentTo(i))
				Expect(value.(*demoStruct).value).Should(BeEquivalentTo(fmt.Sprint(i)))
			}
		})

		It("Given a fibHeap, when call Delete api with a nil value, it should return error.", func() {
			Expect(heap.Delete(nil)).Should(HaveOccurred())
		})

		It("Given a fibHeap inserted multiple values, when call Delete api with a non-exists value, it should return error.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i)
				demo.value = fmt.Sprint(i)
				heap.Insert(demo)
			}
			Expect(heap.Num()).Should(BeEquivalentTo(1000))

			deleteDemo := new(demoStruct)
			deleteDemo.tag = 10000
			deleteDemo.key = float64(10000)
			deleteDemo.value = fmt.Sprint(10000)
			Expect(heap.Delete(deleteDemo)).Should(HaveOccurred())
			Expect(heap.Num()).Should(BeEquivalentTo(1000))
		})

		It("Given a fibHeap inserted multiple values, when call Delete api, it should remove the value from the heap.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i)
				demo.value = fmt.Sprint(i)
				heap.Insert(demo)
			}
			Expect(heap.Num()).Should(BeEquivalentTo(1000))

			for i := 0; i < 1000; i++ {
				deleteDemo := new(demoStruct)
				deleteDemo.tag = i
				deleteDemo.key = float64(i)
				deleteDemo.value = fmt.Sprint(i)
				Expect(heap.Delete(deleteDemo)).ShouldNot(HaveOccurred())
			}
			Expect(heap.Num()).Should(BeEquivalentTo(0))
		})
	})

	Context("union tests", func() {
		BeforeEach(func() {
			heap = NewFibHeap()
			anotherHeap = NewFibHeap()
		})

		AfterEach(func() {
			heap = nil
			anotherHeap = nil
		})

		It("Given two empty fibHeaps, when call Union api, it should return an empty fibHeap.", func() {
			heap.Union(anotherHeap)
			Expect(heap.Minimum()).Should(BeNil())
			Expect(heap.Num()).Should(BeEquivalentTo(0))
		})

		It("Given one empty fibHeap and one non-empty fibHeap, when Union the non-empty one into the empty one, it should retern a new heap with the number and min of the non-empty heap.", func() {
			rand.Seed(time.Now().Unix())
			for i := 0; i < int(rand.Int31n(1000)); i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				anotherHeap.Insert(demo)
			}
			number := anotherHeap.Num()
			min := anotherHeap.Minimum()

			heap.Union(anotherHeap)
			Expect(heap.Minimum()).Should(BeEquivalentTo(min))
			Expect(heap.Num()).Should(BeEquivalentTo(number))
		})

		It("Given one empty fibHeap and one non-empty fibHeap, when Union the empty one into the non-empty one, it should retern a new heap with the number and min of the non-empty heap.", func() {
			rand.Seed(time.Now().Unix())
			for i := 0; i < int(rand.Int31n(1000)); i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}
			number := heap.Num()
			min := heap.Minimum()

			heap.Union(anotherHeap)
			Expect(heap.Minimum()).Should(BeEquivalentTo(min))
			Expect(heap.Num()).Should(BeEquivalentTo(number))
		})

		It("Given two fibHeap with multiple values, when call ExtractMin api after unioned, it should extract the minimum value inserted into both heaps.", func() {
			rand.Seed(time.Now().Unix())
			for i := 0; i < 5000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}
			for i := 0; i < 5000; i++ {
				anotherdemo := new(demoStruct)
				anotherdemo.tag = i + 5000
				anotherdemo.key = rand.Float64()
				anotherdemo.value = fmt.Sprint(anotherdemo.key)
				anotherHeap.Insert(anotherdemo)
			}
			min := heap.Minimum().(*demoStruct).key
			if anotherHeap.Minimum().(*demoStruct).key < min {
				min = anotherHeap.Minimum().(*demoStruct).key
			}
			heap.Union(anotherHeap)

			Expect(heap.Num()).Should(BeEquivalentTo(10000))
			lastKey := heap.Minimum().(*demoStruct).key
			Expect(lastKey).Should(BeEquivalentTo(min))
			for i := 0; i < 10000; i++ {
				extracted := heap.ExtractMin().(*demoStruct)
				Expect(extracted.key).Should(BeNumerically(">=", lastKey))
				Expect(heap.Num()).Should(BeEquivalentTo(9999 - i))
				lastKey = extracted.key
			}
			Expect(heap.Num()).Should(BeEquivalentTo(0))
		})
	})

	Context("index tests", func() {
		BeforeEach(func() {
			heap = NewFibHeap()
			anotherHeap = NewFibHeap()
		})

		AfterEach(func() {
			heap = nil
			anotherHeap = nil
		})

		It("Given one fibHeap, when Insert values with same tag, it should return an error.", func() {
			demo := new(demoStruct)
			demo.tag = 1
			demo.key = 1
			demo.value = "1"
			err := heap.Insert(demo)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(heap.Minimum()).Should(BeEquivalentTo(demo))
			Expect(heap.Num()).Should(BeEquivalentTo(1))
			err = heap.Insert(demo)
			Expect(err).Should(HaveOccurred())
			Expect(heap.Minimum()).Should(BeEquivalentTo(demo))
			Expect(heap.Num()).Should(BeEquivalentTo(1))
		})

		It("Given two fibHeaps which both has value with same tag, when call Union, it should return an error.", func() {
			demo := new(demoStruct)
			demo.tag = 1
			demo.key = 1
			demo.value = "1"
			heap.Insert(demo)
			anotherDemo := new(demoStruct)
			anotherDemo.tag = 1
			anotherDemo.key = 2
			anotherDemo.value = "2"
			anotherHeap.Insert(anotherDemo)

			err := heap.Union(anotherHeap)
			Expect(err).Should(HaveOccurred())
			Expect(heap.Minimum()).Should(BeEquivalentTo(demo))
			Expect(heap.Num()).Should(BeEquivalentTo(1))
			Expect(anotherHeap.Minimum()).Should(BeEquivalentTo(anotherDemo))
			Expect(anotherHeap.Num()).Should(BeEquivalentTo(1))
		})

		It("Given one fibHeaps which has not a value with TAG, when GetTag this TAG, it should return nil.", func() {
			rand.Seed(time.Now().Unix())
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}

			Expect(heap.GetTag(10000)).Should(BeNil())
		})

		It("Given one fibHeaps which has a value with TAG, when GetTag this TAG, it should return the value with this TAG.", func() {
			rand.Seed(time.Now().Unix())
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = rand.Float64()
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}
			tagValue := new(demoStruct)
			tagValue.tag = 10000
			tagValue.key = 10000
			tagValue.value = "10000"
			heap.Insert(tagValue)

			Expect(heap.GetTag(10000)).Should(BeEquivalentTo(tagValue))
			Expect(heap.Num()).Should(BeEquivalentTo(1001))
		})

		It("Given one fibHeaps which has not a value with TAG, when ExtractTag this TAG, it should return nil.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i)
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}
			Expect(heap.Num()).Should(BeEquivalentTo(1000))

			Expect(heap.ExtractTag(1000)).Should(BeNil())
			Expect(heap.Num()).Should(BeEquivalentTo(1000))
		})

		It("Given one fibHeaps which has a value with TAG, when ExtractTag this TAG, it should extract the value with this TAG from the heap.", func() {
			for i := 0; i < 1000; i++ {
				demo := new(demoStruct)
				demo.tag = i
				demo.key = float64(i)
				demo.value = fmt.Sprint(demo.key)
				heap.Insert(demo)
			}
			Expect(heap.Num()).Should(BeEquivalentTo(1000))

			Expect(heap.ExtractTag(999).(*demoStruct).value).Should(BeEquivalentTo(fmt.Sprint(999)))
			Expect(heap.Num()).Should(BeEquivalentTo(999))
			Expect(heap.Minimum().(*demoStruct).value).Should(BeEquivalentTo(fmt.Sprint(0)))
		})
	})

	Context("debug test", func() {
		BeforeEach(func() {
			heap = NewFibHeap()
		})

		AfterEach(func() {
			heap = nil
		})

		It("Given one fibHeaps which some values, when call String api, it should retern the internal debug string.", func() {
			for i := 1; i < 5; i++ {
				for j := 10; j < 15; j++ {
					demo := new(demoStruct)
					demo.tag = i * j
					demo.key = float64(i * j)
					demo.value = fmt.Sprint(demo.key)
					heap.Insert(demo)
				}
				heap.ExtractMin()
			}

			debugMsg := "Total number: 16, Root Size: 1, Index size: 16,\n" +
				"Current minimun: key(14.000000), tag(14), value(&{14 14 14}),\n" +
				"Heap detail:\n" +
				"< 14.000000 < 56.000000 28.000000 < 42.000000 > 30.000000 < 33.000000 36.000000 < 39.000000 > > 20.000000 < 22.000000 24.000000 < 26.000000 > 40.000000 < 44.000000 48.000000 < 52.000000 > > > > > \n"
			Expect(heap.String()).Should(BeEquivalentTo(debugMsg))
		})
	})

	Context("benchmark", func() {
		BeforeEach(func() {
			heap = NewFibHeap()
		})

		AfterEach(func() {
			heap = nil
		})

		Measure("Benchmark Go Fibonacci Heap", func(b Benchmarker) {
			rand.Seed(time.Now().Unix())
			b.Time("1000000 radom operations", func() {
				var (
					insert, minimun, extract, decrease, delete int64
					temp                                       *demoStruct
				)
				for i := 0; i < 1000000; i++ {
					if i%3 == 0 {
						demo := new(demoStruct)
						demo.tag = i
						demo.key = rand.Float64()
						demo.value = fmt.Sprint(demo.key)
						if heap.Insert(demo) == nil {
							insert++
						}
					}
					if i%5 == 0 {
						if heap.ExtractMin() != nil {
							extract++
						}
					}
					if i%7 == 0 {
						min := heap.Minimum()
						if min != nil {
							temp = min.(*demoStruct)
							minimun++
						}
					}
					if i%11 == 0 {
						if temp != nil {
							temp.key = temp.key / 2
							if heap.DecreaseKey(temp) == nil {
								decrease++
							}
						}
					}
					if i%13 == 0 {
						if temp != nil {
							if heap.Delete(temp) == nil {
								delete++
							}
						}
					}
				}
				fmt.Println("Final heap size:", heap.Num())
				fmt.Println("Total insert:", insert, "Total minimun:", minimun, "Total extract:", extract, "Total decrease:", decrease, "Total delete:", delete)
				Expect(heap.Num()).Should(BeEquivalentTo(insert - extract - delete))
			})
		}, 10)
	})
})

type demoStruct struct {
	tag   int
	key   float64
	value string
}

func (demo *demoStruct) Tag() interface{} {
	return demo.tag
}

func (demo *demoStruct) Key() float64 {
	return demo.key
}

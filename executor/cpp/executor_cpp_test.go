package cpp

// func TestArrange(t *testing.T) {
// 	problem := mock.NewMockProblem()
// 	defer problem.RemoveProblemDir()

// 	executorCpp := NewExecutorCpp(
// 		problem,
// 		SourceCodePath{
// 			MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
// 			IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
// 		},
// 	)

// 	err := executorCpp.Arrange()
// 	if err != nil {
// 		t.Errorf("Arrange() failed: %v", err)
// 	}

// 	t.Run("dest.cppがdest_expected.cppと等しい", func(t *testing.T) {
// 		arrangedFile, _ := executorCpp.ArrangedFile()
// 		defer arrangedFile.Close()

// 		destCppPath := arrangedFile.Name()
// 		destCppExpectedPath := fmt.Sprintf("%s/executor/cpp/assets/dest_expected.cpp", setting.RootDir)
// 		//２つのファイルを比較
// 		destCpp, _ := os.ReadFile(destCppPath)
// 		destCppExpected, _ := os.ReadFile(destCppExpectedPath)

// 		if string(destCpp) != string(destCppExpected) {
// 			t.Errorf("arrange() failed: dest.cpp is not same as expected")
// 		}
// 	})
// }

// func TestExecute(t *testing.T) {
// 	problem := mock.NewMockProblem()
// 	defer problem.RemoveProblemDir()

// 	executorCpp := NewExecutorCpp(
// 		problem,
// 		SourceCodePath{
// 			MainPath:       fmt.Sprintf("%s/executor/cpp/assets/main.cpp", setting.RootDir),
// 			IncludeDirPath: fmt.Sprintf("%s/executor/cpp/assets/include", setting.RootDir),
// 		},
// 	)

// 	executorCpp.Arrange()
// 	executorCpp.Compile()

// 	inputFilePath := fmt.Sprintf("%s/executor/cpp/assets/input.txt", setting.RootDir)
// 	outputFilePath := fmt.Sprintf("%s/executor/cpp/assets/output.txt", setting.RootDir)

// 	inputFile, _ := os.Open(inputFilePath)
// 	outputFile, _ := os.Create(outputFilePath)
// 	defer inputFile.Close()
// 	defer outputFile.Close()
// 	defer os.Remove(outputFilePath)

// 	err := executorCpp.Execute(inputFile, outputFile, os.Stderr)
// 	if err != nil {
// 		t.Errorf("Execute() failed: %v", err)
// 	}

// 	t.Run("出力結果が正しい", func(t *testing.T) {
// 		output, _ := os.ReadFile(outputFilePath)
// 		fmt.Println(string(output))
// 		expectedOutput := "6\n"

// 		if string(output) != expectedOutput {
// 			t.Errorf("Execute() failed: output is not same as expected")
// 		}
// 	})
// }

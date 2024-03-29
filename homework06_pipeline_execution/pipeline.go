package homework06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return done
	}
	output := in
	for _, stage := range stages {
		stageCh := make(chan interface{})
		go func(output Out, stageCh Bi, done In) {
			defer close(stageCh)
			for {
				select {
				case <-done:
					return
				case val, ok := <-output:
					if !ok {
						return
					}
					stageCh <- val
				}
			}
		}(output, stageCh, done)
		output = stage(stageCh)
	}

	return output
}

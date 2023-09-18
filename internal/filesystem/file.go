package filesystem

import "os"

func (f *file) OpenFile() (*os.File, error) {
	file, err := os.Open(*f.name)
	if err != nil {
		return nil, err
	}

	return file, err
}

func (f *file) DeleteMany(paths []string) error {
	for _, k := range paths {
		if _, err := os.Stat(k); err != nil {
			return err
		}

		err := os.Remove(k)

		if err != nil {
			return err
		}
	}

	return nil
}

package fns

import (
	"github.com/gographics/imagick/imagick"
	"os"
)

// TransformImage transforms an image.
// See: https://gwdg.de/hpc/_publications/peoospfkdk22/publication.pdf Listing A1.
func TransformImage(inputImage []byte) error {
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(inputImage)
	if err != nil {
		return err
	}

	err = mw.BlurImage(0, 0.5)
	if err != nil {
		return err
	}
	err = mw.AddNoiseImage(imagick.NOISE_MULTIPLICATIVE_GAUSSIAN, 1)
	if err != nil {
		return err
	}
	err = mw.AddNoiseImage(imagick.NOISE_LAPLACIAN, 1)
	if err != nil {
		return err
	}
	err = mw.EnhanceImage()
	if err != nil {
		return err
	}
	err = mw.NegateImage(false)
	if err != nil {
		return err
	}
	err = mw.NormalizeImage()
	if err != nil {
		return err
	}
	// other transformations as mentioned in the paper are not available in the Go bindings

	// create temporary file
	file, err := os.Create("/tmp/scap/results.out")
	if err != nil {
		return err
	}
	defer file.Close()
	fl, err := os.OpenFile(file.Name(), os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer fl.Close()

	err = mw.WriteImageFile(fl)
	if err != nil {
		return err
	}

	return nil
}

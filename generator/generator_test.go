package generator

import "testing"

func TestGenerateIntNumber(t *testing.T) {
	t.Run("min = max", func(t *testing.T) {
		min, max := 5, 5
		num, err := GenerateIntNumber(min, max)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if num != 5 {
			t.Errorf("expected 5, got %d", num)
		}
	})
	t.Run("min < max", func(t *testing.T) {
		min, max := 3, 7
		num, err := GenerateIntNumber(min, max)
		if num < min || num > max {
			t.Errorf("expected %d <= num <= %d, got %d", min, max, num)
		}
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
	t.Run("min > max", func(t *testing.T) {
		min, max := 7, 3
		_, err := GenerateIntNumber(min, max)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

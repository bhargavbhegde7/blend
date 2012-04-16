
package blend

import (
  import "image/color"
)

const (
  MULTIPLY = iota
  SCREEN
  OVERLAY
  SOFT_LIGHT
  HARD_LIGHT
  COLOR_DODGE
  COLOR_BURN
  LINEAR_COLOR_DODGE
  LINEAR_COLOR_BURN
  DARKEN
  LIGHTEN
  DIFFERENCE
  EXCLUSION
  REFLEX
  LINEAR_LIGHT
  PIN_LIGHT
  VIVID_LIGHT
  HARD_MIX
  // Blending modes in HSL color model.
  HUE
  COLOR
  LUMINOSITY
)
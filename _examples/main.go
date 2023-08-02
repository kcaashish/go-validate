package main

import (
	"fmt"

	"github.com/gookit/validate"
)

var data = map[string]any{
	"patient_header": map[string]any{
		"injury_date":    "2023-05-23",
		"patient_status": "Other Disposition",
		"disch_disp":     "OtherDisposition",
		"status":         "COMPLETE",
		"injury_time":    "2023-05-23T23:00:00.000Z",
		"abc":            "123",
		"aba":            nil,
	},
	"coding": []map[string]any{
		{
			"dos": "2023-05-23",
			"details": map[string]any{
				"pro": map[string]any{
					"em": map[string]any{
						"em_level":                      "123",
						"em_modifier1":                  "lorem yus ",
						"em_modifier2":                  "lorem",
						"em_provider":                   "lorem",
						"em_midlevel":                   "lorem",
						"em_type":                       "lorem",
						"em_downcode":                   "true",
						"em_reason":                     "lorem",
						"mdm_total":                     "lorem",
						"qa_disagree":                   "true",
						"qa_note":                       "lorem",
						"mid":                           "lorem",
						"dc_note":                       "lorem",
						"shared":                        "true",
						"charge_code_select":            "lorem",
						"qa_type":                       "lorem",
						"qa_disagree_em_modifier1":      "true",
						"qa_note_em_modifier1":          "lorem",
						"qa_disagree_em_modifier2":      "true",
						"qa_note_em_modifier2":          "lorem",
						"qa_disagree_physician":         "true",
						"qa_note_physician":             "lorem",
						"qa_disagree_shared_visit":      "true",
						"qa_note_shared_visit":          "lorem",
						"qa_disagree_mid_provider":      "true",
						"qa_note_mid_provider":          "lorem",
						"qa_disagree_downcode":          "true",
						"qa_note_downcode":              "lorem",
						"qa_type_em_modifier1":          "lorem",
						"qa_type_em_modifier2":          "lorem",
						"qa_type_physician":             "lorem",
						"qa_type_shared_visit":          "lorem",
						"qa_type_mid_provider":          "lorem",
						"qa_type_downcode":              "lorem",
						"billing_provider":              "lorem",
						"secondary_provider":            "lorem",
						"resident_provider":             "lorem",
						"qa_disagree_resident_provider": "true",
						"qa_note_resident_provider":     "lorem",
						"qa_type_resident_provider":     "lorem",
						"patient_status_id":             1,
						"patient_status_disagree":       "true",
						"qa_note_patient_status":        "lorem",
						"procedure_modifier3":           "lorem",
						"qa_disagree_em_modifier3":      "true",
						"qa_type_em_modifier3":          "lorem",
						"qa_note_em_modifier3":          "lorem",
						"procedure_modifier4":           "lorem",
						"qa_disagree_em_modifier4":      "true",
						"qa_type_em_modifier4":          "lorem",
						"qa_note_em_modifier4":          "lorem",
					},
					"downcode": []map[string]any{
						{
							"em_downcode":          "true",
							"em_reason":            "downcode reason",
							"mdm_total":            "lorem",
							"dc_note":              "lorem",
							"qa_disagree_downcode": "true",
							"qa_note_downcode":     "lorem",
							"qa_type_downcode":     "lorem",
						},
						{
							"em_downcode":          "true",
							"em_reason":            "some other thing",
							"mdm_total":            "lorem",
							"dc_note":              "lorem",
							"qa_disagree_downcode": "true",
							"qa_note_downcode":     "lorem",
							"qa_type_downcode":     "lorem",
						},
					},
					"cpt": []map[string]any{
						{
							"procedure_num":                 "lorem",
							"procedure_qty":                 1,
							"procedure_modifier1":           "lorem",
							"procedure_modifier2":           "lorem",
							"procedure_provider":            "lorem",
							"mid_provider":                  "lorem",
							"qa_disagree":                   "true",
							"qa_note":                       "lorem",
							"procedure_date":                "2023-05-23T00:00:00",
							"charge_code_select":            "lorem",
							"specialist":                    "true",
							"qa_type":                       "lorem",
							"qa_disagree_em_modifier1":      "true",
							"qa_note_em_modifier1":          "lorem",
							"qa_disagree_em_modifier2":      "true",
							"qa_note_em_modifier2":          "lorem",
							"qa_disagree_quantity":          "true",
							"qa_note_quantity":              "lorem",
							"qa_disagree_date_performed":    "true",
							"qa_note_date_performed":        "lorem",
							"qa_disagree_physician":         "true",
							"qa_note_physician":             "lorem",
							"qa_disagree_mid_provider":      "true",
							"qa_note_mid_provider":          "lorem",
							"qa_type_em_modifier1":          "lorem",
							"qa_type_em_modifier2":          "lorem",
							"qa_type_quantity":              "lorem",
							"qa_type_date_performed":        "lorem",
							"qa_type_physician":             "lorem",
							"qa_type_mid_provider":          "lorem",
							"billing_provider":              "lorem",
							"secondary_provider":            "lorem",
							"resident_provider":             "lorem",
							"qa_disagree_resident_provider": "true",
							"qa_note_resident_provider":     "lorem",
							"qa_type_resident_provider":     "lorem",
							"patient_status_id":             1,
							"patient_status_disagree":       "true",
							"qa_note_patient_status":        "lorem",
							"procedure_modifier3":           "lorem",
							"qa_disagree_em_modifier3":      "true",
							"qa_type_em_modifier3":          "lorem",
							"qa_note_em_modifier3":          "lorem",
							"procedure_modifier4":           "lorem",
							"qa_disagree_em_modifier4":      "true",
							"qa_type_em_modifier4":          "lorem",
							"qa_note_em_modifier4":          "lorem",
						},
					},
					"special": []map[string]any{
						{
							"procedure_num":      "lorem",
							"procedure_qty":      1,
							"procedure_provider": "lorem",
							"qa_disagree":        "true",
							"qa_note":            "lorem",
							"procedure_date":     "2023-05-23T00:00:00",
							"qa_type":            "lorem",
						},
					},
					"hcpcs": []map[string]any{
						{
							"hcpcs_num":               "lorem",
							"hcpcs_qty":               1,
							"hcpcs_modifier1":         "lorem",
							"hcpcs_modifier2":         "lorem",
							"hcpcs_provider":          "lorem",
							"qa_disagree":             "true",
							"qa_note":                 "lorem",
							"qa_type":                 "lorem",
							"patient_status_id":       1,
							"patient_status_disagree": "true",
							"qa_note_patient_status":  "lorem",
							"hcpcs_modifier3":         "lorem",
							"hcpcs_modifier4":         "lorem",
							"charge_code_select":      "lorem",
						},
					},
				},
				"fac": map[string]any{
					"em": map[string]any{
						"em_level":                      "lorem",
						"em_modifier1":                  "lorem",
						"em_modifier2":                  "lorem",
						"em_provider":                   "lorem",
						"qa_disagree":                   "true",
						"qa_note":                       "lorem",
						"mid":                           "lorem",
						"charge_code_select":            "lorem",
						"qa_type":                       "lorem",
						"qa_disagree_em_modifier1":      "true",
						"qa_note_em_modifier1":          "lorem",
						"qa_disagree_em_modifier2":      "true",
						"qa_note_em_modifier2":          "lorem",
						"qa_disagree_physician":         "true",
						"qa_note_physician":             "lorem",
						"qa_disagree_mid_provider":      "true",
						"qa_note_mid_provider":          "lorem",
						"qa_type_em_modifier1":          "lorem",
						"qa_type_em_modifier2":          "lorem",
						"qa_type_physician":             "lorem",
						"qa_type_mid_provider":          "lorem",
						"billing_provider":              "lorem",
						"secondary_provider":            "lorem",
						"resident_provider":             "lorem",
						"qa_disagree_resident_provider": "true",
						"qa_note_resident_provider":     "lorem",
						"qa_type_resident_provider":     "lorem",
						"patient_status_id":             1,
						"patient_status_disagree":       "true",
						"qa_note_patient_status":        "lorem",
						"procedure_modifier3":           "lorem",
						"qa_disagree_em_modifier3":      "true",
						"qa_type_em_modifier3":          "lorem",
						"qa_note_em_modifier3":          "lorem",
						"procedure_modifier4":           "lorem",
						"qa_disagree_em_modifier4":      "true",
						"qa_type_em_modifier4":          "lorem",
						"qa_note_em_modifier4":          "lorem",
					},
					"cpt": []map[string]any{
						{
							"procedure_num":                 "lorem",
							"procedure_qty":                 1,
							"procedure_modifier1":           "lorem",
							"procedure_modifier2":           "lorem",
							"procedure_dx":                  "dx.lorem",
							"procedure_provider":            "lorem",
							"mid_provider":                  "lorem",
							"qa_disagree":                   "true",
							"qa_note":                       "lorem",
							"procedure_date":                "2023-05-23T00:00:00",
							"icd10_pcs":                     "icd10_pcs",
							"charge_code_select":            "lorem",
							"qa_type":                       "lorem",
							"qa_disagree_em_modifier1":      "true",
							"qa_note_em_modifier1":          "lorem",
							"qa_disagree_em_modifier2":      "true",
							"qa_note_em_modifier2":          "lorem",
							"qa_disagree_quantity":          "true",
							"qa_note_quantity":              "lorem",
							"qa_disagree_date_performed":    "true",
							"qa_note_date_performed":        "lorem",
							"qa_disagree_physician":         "true",
							"qa_note_physician":             "lorem",
							"qa_disagree_mid_provider":      "true",
							"qa_note_mid_provider":          "lorem",
							"qa_type_em_modifier1":          "lorem",
							"qa_type_em_modifier2":          "lorem",
							"qa_type_quantity":              "lorem",
							"qa_type_date_performed":        "lorem",
							"qa_type_physician":             "lorem",
							"qa_type_mid_provider":          "lorem",
							"qa_disagree_procedure_dx":      "true",
							"qa_note_procedure_dx":          "lorem",
							"qa_type_procedure_dx":          "lorem",
							"qa_disagree_icd10_pcs":         "true",
							"qa_note_icd10_pcs":             "lorem",
							"qa_type_icd10_pcs":             "lorem",
							"billing_provider":              "lorem",
							"secondary_provider":            "lorem",
							"resident_provider":             "lorem",
							"qa_disagree_resident_provider": "true",
							"qa_note_resident_provider":     "lorem",
							"qa_type_resident_provider":     "lorem",
							"patient_status_id":             1,
							"patient_status_disagree":       "true",
							"qa_note_patient_status":        "lorem",
							"procedure_modifier3":           "lorem",
							"qa_disagree_em_modifier3":      "true",
							"qa_type_em_modifier3":          "lorem",
							"qa_note_em_modifier3":          "lorem",
							"procedure_modifier4":           "lorem",
							"qa_disagree_em_modifier4":      "true",
							"qa_type_em_modifier4":          "lorem",
							"qa_note_em_modifier4":          "lorem",
						},
					},
					"special": []map[string]any{
						{
							"procedure_num":      "lorem",
							"procedure_qty":      1,
							"procedure_provider": "lorem",
							"qa_disagree":        "true",
							"qa_note":            "lorem",
							"procedure_date":     "2023-05-23T00:00:00",
							"qa_type":            "lorem",
						},
					},
					"hcpcs": []map[string]any{
						{
							"hcpcs_num":               "lorem",
							"hcpcs_qty":               1,
							"hcpcs_modifier1":         "lorem",
							"hcpcs_modifier2":         "lorem",
							"hcpcs_provider":          "lorem",
							"qa_disagree":             "true",
							"qa_note":                 "lorem",
							"qa_type":                 "lorem",
							"patient_status_id":       1,
							"patient_status_disagree": "true",
							"qa_note_patient_status":  "lorem",
							"hcpcs_modifier3":         "lorem",
							"hcpcs_modifier4":         "lorem",
							"charge_code_select":      "lorem",
						},
					},
				},
				"dx": map[string]any{
					"pro": []map[string]any{
						{
							"code":              "lorem",
							"dx_reason":         "lorem",
							"qa_disagree":       "true",
							"qa_note":           "lorem",
							"qa_type":           "lorem",
							"non_specific_code": "true",
						},
						{
							"code":              "lorem",
							"dx_reason":         "lorem",
							"qa_disagree":       "true",
							"qa_note":           "lorem",
							"qa_type":           "lorem",
							"non_specific_code": "true",
						},
					},
					"fac": []map[string]any{
						{
							"code":              "lorem",
							"dx_reason":         "lorem",
							"qa_disagree":       "true",
							"qa_note":           "lorem",
							"qa_type":           "lorem",
							"non_specific_code": "true",
						},
					},
				},
				"cdi": map[string]any{
					"pro": []map[string]any{
						{
							"enc_cdi_reason": "lorem",
							"cdi_note":       "lorem",
							"qty":            1,
							"mid_level":      "lorem",
							"qa_disagree":    "true",
							"qa_note":        "lorem",
							"qa_type":        "lorem",
						},
					},
					"fac": []map[string]any{
						{
							"enc_cdi_reason": "lorem",
							"cdi_note":       "lorem",
							"qty":            1,
							"mid_level":      "lorem",
							"qa_disagree":    "true",
							"qa_note":        "lorem",
							"qa_type":        "lorem",
						},
					},
				},
				"notes": []map[string]any{
					{
						"coder_notes":  "from coder notes",
						"client_notes": "client notes here",
					},
				},
			},
		},
	},
}

var rules = map[string]string{
	"patient_header.injury_time":                   "if_not_null:gte_field:patient_header.injury_date",
	"patient_header.injury_date":                   "required",
	"patient_header.abc":                           "required",
	"patient_header.aba":                           "gte_field:patient_header.abc",
	"coding":                                       "required",
	"coding.#.dos":                                 "required",
	"coding.#.details.dx.pro":                      "required",
	"coding.#.details.dx.pro.#.code":               "required",
	"coding.#.details.pro.cpt.#.procedure_num":     "required",
	"coding.#.details.pro.cpt.#.procedure_qty":     "required",
	"coding.#.details.pro.cpt.#.specialist":        "required",
	"coding.#.details.pro.special.#.procedure_num": "required",
	"coding.#.details.pro.special.#.procedure_qty": "required",
	"coding.#.details.pro.hcpcs.#.hcpcs_num":       "required",
	"coding.#.details.pro.hcpcs.#.hcpcs_qty":       "required",
	"coding.#.details.fac.em.em_level":             "required",
	"coding.#.details.fac.em.billing_provider":     "required",
	"coding.#.details.fac.cpt.#.procedure_num":     "required",
	"coding.#.details.fac.cpt.#.procedure_qty":     "required",
	"coding.#.details.fac.special.#.procedure_num": "required",
	"coding.#.details.fac.special.#.procedure_qty": "required",
	"coding.#.details.fac.hcpcs.#.hcpcs_num":       "required",
	"coding.#.details.fac.hcpcs.#.hcpcs_qty":       "required",
	"coding.#.details.dx.fac.#.code":               "required",
	"coding.#.details.cdi.pro.#.enc_cdi_reason":    "required",
	"coding.#.details.cdi.pro.#.qty":               "required",
	"coding.#.details.cdi.fac.#.enc_cdi_reason":    "required",
	"coding.#.details.cdi.fac.#.qty":               "required",
	"coding.#.details.pro.em":                      "required_with:coding.#.details.pro",
	"coding.#.details.pro.em.em_level":             "required_with:coding.#.details.pro.em",
	"coding.#.details.pro.em.em_downcode":          "required_with:coding.#.details.pro.em",
	"coding.#.details.pro.em.shared":               "required_with:coding.#.details.pro.em",
	"coding.#.details.pro.em.billing_provider":     "required_with:coding.#.details.pro.em",
	"coding.#.details.pro.downcode.#.em_downcode":  "required_with:coding.#.details.pro.downcode",
	"coding.#.details.pro.downcode.#.em_reason":    "required_with:coding.#.details.pro.downcode",
	"coding.#.details.pro.downcode.#.mdm_total":    "required_with:coding.#.details.pro.downcode",
}

func main() {
	validator := validate.Map(data)
	for field, rule := range rules {
		validator.StringRule(field, rule)
	}
	fmt.Println(validator.Validate(), validator.ValidateE())
}
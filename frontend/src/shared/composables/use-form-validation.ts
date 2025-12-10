import * as v from 'valibot'
import { ref } from 'vue'

export const useFormValidation = <TSchema extends v.GenericSchema, TForm = v.InferInput<TSchema>>(
  schema: TSchema,
) => {
  type ErrorsType = Partial<Record<keyof TForm, string>>

  const errors = ref<ErrorsType>({})

  const validate = (form: TForm): boolean => {
    errors.value = {}

    const result = v.safeParse(schema, form)

    if (!result.success) {
      for (const issue of result.issues) {
        const field = issue.path?.[0]?.key as keyof TForm
        if (field) {
          errors.value[field] = issue.message
        }
      }
      return false
    }

    return true
  }

  const clearErrors = () => {
    errors.value = {}
  }

  return {
    errors,
    validate,
    clearErrors,
  }
}

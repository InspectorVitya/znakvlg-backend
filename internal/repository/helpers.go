package repository

const selectInfoCompanies = `
with plates_wait_audit as (
    select company_id, sum(count_plate) as wait_audit from products where status_product_id = 1 group by company_id
), plates_wait_audit_today as (
    select company_id, sum(count_plate) as wait_audit_today
    from products where status_product_id = 1
                    and (products.created_at >= $1 and products.created_at <= $2) group by company_id
), plates_audit as (
    select company_id, sum(count_plate) as audit
    from products where status_product_id = 2 group by company_id
), plates_audit_today as (
    select company_id, sum(count_plate) as audit_today
    from products where status_product_id = 2
                    and (products.created_at >= $1 and products.created_at <= $2) group by company_id
), deffect as (
    select company_id, sum(count_plate) as deffect
    from defect_products as deffect group by company_id
), deffect_today as (
    select company_id, sum(count_plate) as deffect_today
    from defect_products as deffect where deffect.created_at >= $1 and deffect.created_at <= $2 group by company_id
), plates_wait_check as (
    select company_id, count(count_plate) as wait_check
    from products where status_product_id = 1 group by company_id
), plates_wait_recheck as (
    select company_id, count(count_plate) as wait_recheck
    from products where status_product_id = 4 group by company_id
), plate_used as (
    select company_id, plate_type_id, sum(count_plate) as plate_used
    from products group by company_id,plate_type_id
),  plate_used_today as (
    select company_id, plate_type_id, sum(count_plate) as plate_used_today
    from products
    where products.created_at >= $1 and products.created_at <= $2
        group by company_id, plate_type_id
)


SELECT c.id, c.title, c.address, c.comment, pt.title as title_plate, s.count,
       COALESCE(plates_wait_audit.wait_audit, 0) as wait_audit,
       COALESCE(plates_wait_audit_today.wait_audit_today, 0) as wait_audit_today,
       COALESCE(plates_audit.audit, 0) as audit,
       COALESCE(plates_audit_today.audit_today, 0) as audit_today,
       COALESCE(deffect.deffect, 0) as deffect,
       COALESCE(deffect_today.deffect_today, 0) as deffect_today,
       COALESCE(plates_wait_check.wait_check, 0) as wait_check,
       COALESCE(plates_wait_recheck.wait_recheck, 0) as wait_recheck,
       COALESCE(plate_used.plate_used, 0) as plate_used,
       COALESCE(plate_used_today.plate_used_today, 0) as plate_used_today
    FROM companies as c
        LEFT join plates_wait_audit on c.id = plates_wait_audit.company_id
        LEFT join plates_wait_audit_today on c.id = plates_wait_audit_today.company_id
        LEFT join plates_audit on c.id = plates_audit.company_id
        LEFT join plates_audit_today on c.id = plates_audit_today.company_id
        LEFT join deffect on c.id = deffect.company_id
        LEFT join deffect_today on c.id = deffect_today.company_id
        LEFT join plates_wait_check on c.id = plates_wait_check.company_id
        LEFT join plates_wait_recheck on c.id = plates_wait_recheck.company_id
        LEFT JOIN store s on c.id = s.company_id
        LEFT JOIN plate_used on c.id = plate_used.company_id and s.plate_type_id = plate_used.plate_type_id
        LEFT JOIN plate_used_today on c.id = plate_used_today.company_id and s.plate_type_id = plate_used_today.plate_type_id
		left join plate_types pt on s.plate_type_id = pt.id;
`


-- Update the total cost of all invoices
UPDATE invoices
SET total = subquery.total_cost
FROM (
         SELECT
             i.id AS invoice_id,
             COALESCE(SUM(ii.quantity * ip.value), 0) AS total_cost
         FROM invoices i
                  LEFT JOIN invoices_items ii ON i.id = ii.invoice_id
                  LEFT JOIN item_prices ip
                            ON ii.item_id = ip.item_id
                                AND current_date BETWEEN ip.begin_date AND ip.end_date
         GROUP BY i.id
     ) AS subquery
WHERE invoices.id = subquery.invoice_id;

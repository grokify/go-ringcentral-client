#!ruby

require 'multi_json'

doc = '
    {
      "id":"5cde163bd6cb00d3034e2332",
      "created_at":"2019-05-17T02:02:35Z",
      "updated_at":"2019-05-17T02:31:03Z",
      "category_ids":[

      ],
      "email":"embbnux.ji@ringcentral.com",
      "enabled":true,
      "external_id":null,
      "firstname":"Nux",
      "gender":null,
      "identity_ids":[

      ],
      "lastname":"Ji",
      "locale":"en",
      "nickname":null,
      "rc_user_id":null,
      "role_id":"5c8bf22b14bf8a84c44c7750",
      "spoken_languages":[
        "en"
      ],
      "team_ids":[

      ],
      "timezone":"",
      "invitation_pending":false
    }
'
doc = '
  {
    "id":"5c8bf22b14bf8a84c44c7753",
    "created_at":"2019-03-15T18:42:51Z",
    "updated_at":"2019-03-15T18:42:51Z",
    "label":"Agent",
    "approve_content":false,
    "assign_intervention":false,
    "author_block_content":false,
    "admin_stamp_answer":true,
    "close_content_thread":true,
    "mute_content":false,
    "update_own_intervention":true,
    "update_intervention":false,
    "publish_content":true,
    "delay_export_content":true,
    "receive_tasks":true,
    "search_contents":true,
    "open_content_thread":true,
    "impersonate_user":false,
    "delete_content_thread":false,
    "use_emoji":true,
    "access_pull_mode":true,
    "read_event":true,
    "read_presence":true,
    "read_identity":true,
    "access_previous_messages":true,
    "use_cobrowsing":true,
    "invite_user":false,
    "create_user":false,
    "read_user":true,
    "update_user":false,
    "manage_users_of_my_teams":false,
    "manage_identities":false,
    "update_identity":true,
    "manage_teams":false,
    "manage_roles":false,
    "manage_own_notifications":false,
    "manage_categories":false,
    "manage_folders":false,
    "manage_custom_notifications":false,
    "read_community":false,
    "create_community":false,
    "update_community":false,
    "read_content_source":false,
    "create_content_source":false,
    "update_content_source":false,
    "manage_chat":false,
    "manage_messaging":false,
    "manage_topologies":false,
    "read_export":false,
    "search_event":false,
    "update_settings":false,
    "manage_tags":false,
    "manage_custom_fields":false,
    "manage_emails_templates":false,
    "manage_api_access_tokens":false,
    "create_and_destroy_extension":false,
    "update_extension":false,
    "update_time_sheet":false,
    "manage_app_sdk_applications":false,
    "access_help_center":true,
    "manage_ice":false,
    "export_identity":false,
    "anonymize_identity":false,
    "lock_identity":false,
    "read_stats":false,
    "read_own_stats":true,
    "monitor_tasks":false,
    "monitor_team_tasks":false,
    "manage_reply_assistant":false,
    "reply_with_assistant":true,
    "manage_rules_engine_rules":false
  }
'

doc = '
{  "id": "5b87a6b2f042de5f94fabf8d",

  "created_at": "2018-08-30T08:11:30Z",
  "updated_at": "2018-08-30T08:11:30Z",
  "content_type": "image/jpeg",
  "filename": "cat.jpeg",
  "foreign_id": null,
  

  "size": 700754,
  "url": "http://domain-test.engagement.dimelo.dev/attachments/5b87a6b2f042de5f94fabf8d",
 

  "video_metadata": [],
  "embed": false,
  "public?": true

}
'
doc = '
{
  "id": "60944e5702bdafb74ec96142",
  "parent_id": "eb3c62690ec9025845fd3495",
  "name": "Technical",
  "created_at": "2012-05-23T01:12:49Z",
  "updated_at": "2012-05-23T01:12:49Z",
  "color": 0,
  "mandatory": false,
  "multiple": true,
  "post_qualification": false,
  "source_ids": [],
  "unselectable": false  }'

doc = '{
  "id":"5c8bf22b14bf8a84c44c773b",
  "created_at":"2019-03-15T18:42:51Z",
  "updated_at":"2019-07-24T09:26:42Z",
  "activity_presence_threshold":45,
  "activity_tracking":true,
  "beginning_of_week":"monday",
  "content_languages":[
    "en"
  ],
  "deny_iframe_integration":false,
  "disable_password_autocomplete":false,
  "dump_in_preprod":false,
  "expire_password_after":7776000,
  "expire_password_enabled":false,
  "fte_duration":7,
  "identity_merge":true,
  "intervention_defer_rates":[
    1800,
    86400,
    259200
  ],
  "intervention_defer_threshold":86400,
  "intervention_rates":[
    3600,
    10800,
    86400,
    259200
  ],
  "locale":"en",
  "multi_lang":false,
  "name":"rc-platform",
  "password_archivable_enabled":false,
  "password_archivable_size":5,
  "password_min_length":6,
  "password_non_word":false,
  "password_numbers":false,
  "password_recovery_disabled":false,
  "push_enabled":true,
  "reply_as_any_identity":false,
  "rtl_support":false,
  "self_approval_required":false,
  "session_timeout":240,
  "sharding_key":"rc_platform",
  "spellchecking":false,
  "style":"ringcentral",
  "third_party_services_disabled":false,
  "timezone":"Pacific Time (US & Canada)",
  "type":"demo",
  "urgent_task_threshold":60,
  "browser_notifications_disabled":false,
  "display_only_unknown_bbcode":false,
  "intervention_closing_period":432000,
  "use_two_letters_avatars":true
}'

doc = '{"id": "73f1cb2938229d7fa222d096",
"source_id": "d19c81948c137d86dac77216",
"source_url": "http://domain-test.answers.dimelo.com/questions/42",
"source_type": "answers",
"thread_id": "26c56bc5b71c5193b6f8c656",
"in_reply_to_id": "58bc74bc920026b30196fdf4",
"in_reply_to_author_id": "57ea9a7f13047d506e65289d",
"author_id": "4f0aa52d656a3d75867f784c",
"creator_id": "ac24dc966bc7ecb74017c0cd",
"foreign_id": "7789",
"type": "question",
"created_from": "synchronizer",
"private_message": false,
"status": "replied",
"intervention_id": "7f946431b6eebffafae642cc",
"language": "fr",
"body": "Hello,\n\nHow to unlock my nokia 3210?\n\nThanks!",
"body_formatted": {
  "text": "Hello,\n\nHow to unlock my nokia 3210?\n\nThanks!",
  "html": "<p>Hello,</p>\n\n<p>How to unlock my nokia 3210?</p>\n\n<p>Thanks!</p>"
 },
"body_input_format": "text",
"title": "Nokia 3210 unlocking",
"attachments_count": 1, "attachments": [{
"id":"5464b5c04d61639684110000", "created_at":"2011-05-05T22:00:00Z",
 "updated_at":"2011-05-05T22:00:00Z", 
 "content_type": "application/pdf", "size": 174784,
"filename":"sso.pdf",
"foreign_id":"123",
"embed":"false",
"public?":"true",
"url":"h​ ttp://domain-test.engagement.dimelo.dev/attachments/5464b5c04d61639684110000​"
}],
"synchronization_status": "success",
"category_ids": [
  "4d0fb475b242228033cbdf6d", "60944e5702bdafb74ac96141"],
   "created_at": "2012-05-24T04:00:44Z",
"updated_at": "2012-05-24T04:00:44Z",
"approval_required": false,
"remotely_deleted": false,
"published": true,
"context_data": {
"customer_id": 1214 
}
}
'
properties = {}

object = MultiJson.load doc

object.each do |name,v|
  prop = {
    type: 'string'
  }
  if v =~ /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$/
    prop[:format] = 'date-time'
  end

  if !!v == v
    prop[:type] = 'boolean'
  elsif v.is_a? Integer
    prop[:type] = 'integer'
  end
  if v.kind_of?(Array)
    prop[:type] = 'array'
    prop[:items] = { type: "string" }
  end

  properties[name] = prop
end

puts doc

puts MultiJson.dump properties, pretty: true